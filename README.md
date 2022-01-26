# Calculator-in-Go-and-Fyne



https://user-images.githubusercontent.com/96500508/151125860-89e04979-d6f2-4665-9c8b-1dec45355e87.mp4



## Prerequisites

Fyne requires 3 basic elements to be present, the Go tools, a C compiler (to connect with system graphics drivers) and an system graphics driver.

1. Download [Go](https://go.dev/dl/)

2. Install C compiler for windows: [TDM-GCC](https://jmeubank.github.io/tdm-gcc/download/)

3. In Windows your graphics driver will already be installed, but it is recommended to ensure they are up to date.

## Downloading

When using Go modules, you will need to set up the module before you can use the package.

`$ cd myapp`

`$ go mod init MODULE_NAME`

You now need to download the Fyne module. This will be done using the following command:

`$ go get fyne.io/fyne/v2`

If you are unsure of how Go modules work, consider reading [Tutorial: Create a Go module](https://go.dev/doc/tutorial/create-module)

