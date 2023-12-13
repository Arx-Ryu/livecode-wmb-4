package main

import (
	"bufio"
	"fmt"
	"livecode-4/model"
	"livecode-4/repository"
	"livecode-4/usecase"
	"math/rand"
	"os"
	"strings"
)

func main () {
	wmb()
}

func wmb () {	
	tbl := repository.NewTableRepo(30)
	menu := repository.NewMenuRepo()
	trx := repository.NewTrxRepo()	
	cus := repository.NewCustomerRepo()	
	log := repository.NewLogRepo()

	testCases(trx, tbl, cus) //Nyalakan Jika Ingin Ada Contoh Data Masuk Dahulu
	
	pilihan := 9
	for pilihan > 0 {
		fmt.Println(strings.Repeat("=",50))
		fmt.Println(`Kasir: 
	1 - Table Status	4 - Print Bill
	2 - Menu Warung Makan	5 - Payment
	3 - Order		6 - Export Log Data
		
	0 - Exit`)
		fmt.Print("Pilihan: ")
		fmt.Scanln(&pilihan)
		switch {
		case pilihan == 1:
			tablePanel(tbl)
			LogData(log, "Open Table Status")
		case pilihan == 2:
			menuPanel(menu, log)
		case pilihan == 3:
			NewTrx(trx, tbl, menu, cus, log)	
		case pilihan == 4:
			PrintTrx(trx, tbl, log)
		case pilihan == 5:
			PayTrx(trx, tbl, log)
		case pilihan == 6:
			ExportData(trx, tbl, menu, cus, log)
		case pilihan == 0:
			fmt.Print("Goodbye")
		default:
			fmt.Printf("Tidak ada Pilihan dengan Nomor %d", pilihan)
		}
	}	
}

func testCases (trx repository.TransRepo, tbl repository.TableRepo, cus repository.CustomerRepo) {
	orderUseCase := usecase.NewCustomerOrderUseCase(trx, tbl)	
	customerUseCase := usecase.NewCustomerUseCase(cus)
	
	var customer model.Customer
	var customerOrder []model.CustomerOrder
	var order model.CustomerOrder
	var selectedMenu model.Menu	
	//Test Cases for Customer Data
	customerName := []string{
		"Kemal Roushdy Jenie", 
		"Samuel Maynard", 
		"Elia Samuel", 
		"Fajar Setiawan", 
		"Niqma Rozalia", 
		"Fira Faiza Tertia", 
		"Irsalina Layalia S", 
		"Vidya Aisya",
	}	
	tableNum := []string{
		"T29",
		"T15",
		"T22",
		"T05",
		"T11",
		"T18",
		"T24",
		"T08",
	}
	menu01 := model.Menu{
		MenuId: "M001",
		MenuName: "Nasi Goreng",
		MenuPrice: 18000,
	}
	menu02 := model.Menu{
		MenuId: "M002",
		MenuName: "Nasi Telor",
		MenuPrice: 12000,
	}
	menu03 := model.Menu{
		MenuId: "M003",
		MenuName: "Nasi Ayam",
		MenuPrice: 16000,
	}
	menu04 := model.Menu{
		MenuId: "M004",
		MenuName: "Indomie",
		MenuPrice: 8000,
	}
	menu05 := model.Menu{
		MenuId: "M005",
		MenuName: "Teh Tawar",
		MenuPrice: 3000,
	}
	menu06 := model.Menu{
		MenuId: "M006",
		MenuName: "Kopi",
		MenuPrice: 4000,
	}
	for i, name := range customerName {
		customerOrder = nil
		for x := 0; x < 3; x++ {
			y := rand.Int()
			if y%2 == 0 {
				if y%3 == 0 { selectedMenu = menu01
				} else { selectedMenu = menu02
				} 
				order = model.CustomerOrder{
					MenuName: selectedMenu,
					Qty: (y%13)+1,
				}
				customerOrder = append(customerOrder, order)
				selectedMenu = menu05
				order = model.CustomerOrder{
					MenuName: selectedMenu,
					Qty: (y%13)+1,
				}
				customerOrder = append(customerOrder, order)	
			} else {
				if y%3 == 0 { selectedMenu = menu03
				} else { selectedMenu = menu04
				}
				order = model.CustomerOrder{
					MenuName: selectedMenu,
					Qty: (y%13)+1,
				}
				customerOrder = append(customerOrder, order)
				selectedMenu = menu06
				order = model.CustomerOrder{
					MenuName: selectedMenu,
					Qty: (y%13)+1,
				}
				customerOrder = append(customerOrder, order)
			}		
		}
		if i < 3 {
			customer = customerUseCase.NewMember(name, true) //Member
		} else {
			customer = customerUseCase.NewMember(name, false) //Non-Member
		}			
		orderUseCase.TakeOrder(customer, tableNum[i], customerOrder)
	}
	fmt.Println("8 Test Cases Succesfully Inputted")
}

func tablePanel (tbl repository.TableRepo) {
	tableUseCase := usecase.NewTableViewUseCase(tbl)	
	tableUseCase.ViewTable()	
}

func menuPanel (menu repository.MenuRepo, log repository.LogRepo) {
	menuUseCase := usecase.NewFindMenuUseCase(menu)

	var pilihan int
	var menuSearch string

	scanner := bufio.NewReader(os.Stdin)

	fmt.Println(strings.Repeat("=",50))
	fmt.Println(`Menu Panel: 
	1 - Show Full Menu
	2 - Search By ID
	3 - Search By Name
	0 - Exit`)
	fmt.Print("Pilihan: ")
	fmt.Scanln(&pilihan)
	switch {
	case pilihan == 1:
		menuUseCase.ShowMenu()
		LogData(log, "Open Full Menu")
	case pilihan == 2:
		fmt.Print("Menu ID to look for: ")
		fmt.Scanln(&menuSearch)
		menuUseCase.FindMenuById(menuSearch)
		LogData(log, "Menu Search By ID")
	case pilihan == 3:
		fmt.Print("Menu Name to look for: ")
		menuSearch, _ = scanner.ReadString('\n')
		menuSearch = strings.TrimRight(menuSearch, "\r\n")
		menuUseCase.FindMenuByName(menuSearch)
		LogData(log, "Menu Search By Name")
	case pilihan == 0:
		fmt.Print("Goodbye")
		LogData(log, "Exit Menu Panel")
	default:
		fmt.Printf("Tidak ada Pilihan dengan Nomor %d", pilihan)
		LogData(log, "Exit Menu Panel")
	}	
}

func NewTrx (trx repository.TransRepo, tbl repository.TableRepo, menu repository.MenuRepo, cus repository.CustomerRepo, log repository.LogRepo) {
	tableUseCase := usecase.NewTableViewUseCase(tbl)
	menuUseCase := usecase.NewFindMenuUseCase(menu)
	orderUseCase := usecase.NewCustomerOrderUseCase(trx, tbl)	
	customerUseCase := usecase.NewCustomerUseCase(cus)

	scanner := bufio.NewReader(os.Stdin)
	
	menuSelect := 1
	var menuId, customerName, tableNum, customerMember string
	var table model.Table
	var customer model.Customer
	var customerOrder []model.CustomerOrder
	var order model.CustomerOrder
	var selectedMenu model.Menu
	var menuQty int

	for menuSelect > 0 {
		fmt.Print("Menu ID to add (0 Exit, 1 Open Menu): ")
		fmt.Scanln(&menuId)
		if menuId == "0" {
			break
		} else if menuId == "1" {
			menuUseCase.ShowMenu()
		} else {
			selectedMenu = menuUseCase.FindMenuById(menuId)
			if selectedMenu.MenuId != "" {
				fmt.Print("Item Qty: ")
				fmt.Scanln(&menuQty)
				order = model.CustomerOrder{
					MenuName: selectedMenu,
					Qty: menuQty,
				}
				customerOrder = append(customerOrder, order)
			} 
		}					
	}
	if len(customerOrder) != 0 {
		fmt.Print("Customer Name: ")
		customerName, _ = scanner.ReadString('\n')
		customerName = strings.TrimRight(customerName, "\r\n")
		customer = customerUseCase.FindByName(customerName)
		if customer.CustomerName == "" {
			fmt.Print("Would you like to register as a member (y/n)?")
			fmt.Scanln(&customerMember)
			if strings.ToLower(customerMember) == "y" {
				customerUseCase.NewMember(customerName, true)
				logBody := "New customer " + customerName + " registered as member"
				LogData(log, logBody)
			} else {
				customerUseCase.NewMember(customerName, false)
				logBody := "New customer " + customerName + " registered"
				LogData(log, logBody)
			}		
		} else if customer.IsMember {
			fmt.Printf("Welcome Back Member %s!\n", customer.CustomerName)
		} else {
			fmt.Printf("Welcome Back %s!\n", customer.CustomerName)
		}
		fmt.Print("Table Reservation (0 if none): ")
		fmt.Scanln(&tableNum)
		if tableNum == "0" {
			table = tableUseCase.GetEmptyTable()
		} else {
			table = tableUseCase.GetReserveTable(tableNum)
			if table.TableNum == "" || !table.TableStatus { //Meja Reservasi Tidak Ketemu atau Dipakai
				fmt.Println("Getting New Empty Table")
				table = tableUseCase.GetEmptyTable()
			} 				
		}
		if table.TableNum == "" { //Tidak ada meja kosong
			logBody := "Placed Order but no empty tables"
			LogData(log, logBody)
			return
		} else {
			tableNum = table.TableNum
			trxId := orderUseCase.TakeOrder(customer, tableNum, customerOrder)
			logBody := "Placed Order " + trxId + " for customer " + customerName + " at table " + tableNum
			LogData(log, logBody)
		}
	} else {
		LogData(log, "Placed Order but Canceled")
	}
}

func PrintTrx (trx repository.TransRepo, tbl repository.TableRepo, log repository.LogRepo) {
	paymentUseCase := usecase.NewCustomerPaymentUseCase(trx, tbl)

	var trxId string

	fmt.Print("Order ID to Print: ")
	fmt.Scanln(&trxId)
	paymentUseCase.PrintTrx(trxId)
	logBody := "Print Bill Order ID " + trxId
	LogData(log, logBody)
}

func PayTrx (trx repository.TransRepo, tbl repository.TableRepo, log repository.LogRepo) {
	paymentUseCase := usecase.NewCustomerPaymentUseCase(trx, tbl)

	var trxId string

	fmt.Print("Order ID to Pay: ")
	fmt.Scanln(&trxId)
	paymentUseCase.Payment(trxId)
	logBody := "Payment Order ID " + trxId
	LogData(log, logBody)
}

func ExportData (trx repository.TransRepo, tbl repository.TableRepo, menu repository.MenuRepo, cus repository.CustomerRepo, log repository.LogRepo) {
	tableUseCase := usecase.NewTableViewUseCase(tbl)
	menuUseCase := usecase.NewFindMenuUseCase(menu)
	orderUseCase := usecase.NewCustomerOrderUseCase(trx, tbl)	
	customerUseCase := usecase.NewCustomerUseCase(cus)
	file := usecase.NewFileUseCase(log)	
	
	var fileLocation, data string
	scanner := bufio.NewReader(os.Stdin)

	fmt.Print("Location to export file: ")
	fileLocation, _ = scanner.ReadString('\n')
	fileLocation = strings.TrimRight(fileLocation, "\r\n") 
	// if !strings.Contains(fileLocation, ".") {
	// 	fmt.Println("This location has no file format (i.e. .txt)")
	// } else {
		
	// }	
	tableData := tableUseCase.ExportData()
		data = fmt.Sprint(strings.Repeat("=", 100)) + fmt.Sprintf("\nTABLE DATA\n%v\n", tableData)
		menuData := menuUseCase.ExportData()
		data = data + fmt.Sprint(strings.Repeat("=", 100)) + fmt.Sprintf("\nMENU DATA\n%v\n", menuData)
		customerData := customerUseCase.ExportData()
		data = data + fmt.Sprint(strings.Repeat("=", 100)) + fmt.Sprintf("\nCUSTOMER DATA\n%v\n", customerData)
		trxData := orderUseCase.ExportData()
		data = data + fmt.Sprint(strings.Repeat("=", 100)) + fmt.Sprintf("\nTRANSACTION DATA\n%v\n", trxData)
		LogData(log, data)		
		output := file.ExportLogData(fileLocation)
		fmt.Print(output)
}

func LogData (log repository.LogRepo, item string) {
	file := usecase.NewFileUseCase(log)	
	file.NewLog(item)	
}