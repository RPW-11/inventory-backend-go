package usecase

import (
	"time"

	"github.com/RPW-11/inventory_management_be/domain"
	"github.com/google/uuid"
)

type inventoryUsecase struct {
	inventoryRepository domain.InventoryRepository
	productRepository   domain.ProductRepository
	warehouseRepository domain.WarehouseRepository
}

func NewInventoryUsecase(ir domain.InventoryRepository, pr domain.ProductRepository, wr domain.WarehouseRepository) domain.InventoryUsecase {
	return &inventoryUsecase{
		inventoryRepository: ir,
		productRepository:   pr,
		warehouseRepository: wr,
	}
}

func (iu *inventoryUsecase) CreateProductInventory(product *domain.Product, warehouses []domain.WarehouseQuantity) (string, error) {
	// check if all warehouses exists
	for _, w := range warehouses {
		_, err := iu.warehouseRepository.GetByID(w.WarehouseID)
		if err != nil {
			return "", err
		}
	}

	// check if the product exists
	_, err := iu.productRepository.GetByID(product.ID)
	if err != nil {
		if err.Error() != "no existing product" {
			return "", err
		}
		// create a product
		product.ID = uuid.NewString()
		iu.productRepository.Create(product)
	}

	// insert product to the inventories
	for _, w := range warehouses {
		// check if the inventory exists
		existingInventory, err := iu.inventoryRepository.GetByProductWarehouseID(product.ID, w.WarehouseID)
		if err != nil {
			if err.Error() != "no existing inventory" {
				return "", err
			}
			// create the inventory
			inventory := domain.Inventory{
				Quantity:    w.ProductQuantity,
				ProductId:   product.ID,
				WarehouseId: w.WarehouseID,
			}

			err = iu.inventoryRepository.Create(&inventory)
			if err != nil {
				return "", err
			}
		} else {
			existingInventory.Quantity += w.ProductQuantity
			newInventory := domain.Inventory{
				ProductId:   existingInventory.ProductId,
				WarehouseId: existingInventory.WarehouseId,
				Quantity:    existingInventory.Quantity,
				UpdatedAt:   time.Now(),
			}
			err = iu.inventoryRepository.ModifyByID(existingInventory.ID, &newInventory)
			if err != nil {
				return "", err
			}
		}
	}

	return product.ID, nil
}

func (iu *inventoryUsecase) GetByID(id int) (domain.Inventory, error) {
	return iu.inventoryRepository.GetByID(id)
}

func (iu *inventoryUsecase) GetProductDetails() ([]domain.ProductDetail, error) {
	productDetails := []domain.ProductDetail{}

	products, err := iu.productRepository.Fetch("")
	if err != nil {
		return productDetails, err
	}

	for _, p := range products {
		ivs, err := iu.inventoryRepository.GetByProductID(p.ID)
		if err != nil {
			return productDetails, err
		}
		productDetail := domain.ProductDetail{
			Product:     p,
			Inventories: []domain.InventoryDetail{},
			ImageUrls:   []string{},
		}

		// grab the quantity and query all the warehouse detail
		for _, i := range ivs {
			warehouse, err := iu.warehouseRepository.GetByID(i.WarehouseId)
			if err != nil {
				return productDetails, err
			}
			productDetail.Inventories = append(productDetail.Inventories, domain.InventoryDetail{
				ID:               i.ID,
				WarehouseID:      i.WarehouseId,
				WarehouseName:    warehouse.Name,
				WarehouseAddress: warehouse.Address,
				ProductQuantity:  i.Quantity,
			})
		}

		// grab all of the images
		images, err := iu.productRepository.GetImagesByProductId(p.ID)
		if err != nil {
			return productDetails, err
		}

		for _, img := range images {
			productDetail.ImageUrls = append(productDetail.ImageUrls, img.ImageUrl)
		}

		// append the product detail to the array
		productDetails = append(productDetails, productDetail)
	}
	return productDetails, err
}

func (iu *inventoryUsecase) ModifyByID(inventoryID int, inventory *domain.Inventory) error {
	return iu.inventoryRepository.ModifyByID(inventoryID, inventory)
}
