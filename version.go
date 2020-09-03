package aldoutil

const (
	version string = "1.3.0"
	// Added field ean on store product.

	// version string = "1.2.0"
	// // Added field storeProductQtd on store product.

	// version string = "1.1.0"
	// // Using new data base (removed fields "new", "changed" and "removed", added fields "removedAt", "StatusCleanedAt").

	// version string = "1.0.5"
	// // Bugfix - 'removed' status has prority over 'unavaiable'.

	// version string = "1.0.4"
	// // Included options on status (unavailable, removed).

	// version string = "1.0.3"
	// // Bugfix - Not show status product as new for products created at zunka site.

	// version string = "1.0.2"
	// // Bugfix - Product status method.

	// version string = "1.0.1"
	// // Product have a method to define status.

	// // Remove sql code.
	// version string = "0.6.0"
)

/*
:: v0.5.0
Included StoreProduct.DealerProductImagesLink struct field.

:: v0.4.2
Removed limit 10 from select all.
Removed unit and multiple from product struct.
*/
