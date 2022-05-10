package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello Welcome")
	var card string
	flag := true
	forFlag := true
	atmBalance, dailyLimit := 20000, 10000
	savingAccBalance, currentAccBalance := 50000, 5000
	for forFlag == true {
		fmt.Println("Please Insert the card")
		fmt.Scanln(&card)
		for flag == true {
			if atmBalance != 0 {
				if len(card) == 12 {
					if strings.HasPrefix(card, "1234") {
						fmt.Println("Enter Pin")
						var pin int
						fmt.Scanln(&pin)
						pinlen := strconv.Itoa(pin)
						if len(pinlen) == 4 {
							if pin == 9874 {
								typeFlag := true
								for typeFlag == true {
									fmt.Printf("Choose Account type \n 1) SAVINGS ACCOUNT 2)CURRENT ACCOUNT\n")
									var accType int
									fmt.Scanln(&accType)
									if accType == 1 || accType == 2 {
										typeFlag = false

										switch accType {
										case 1:
											switchFlag1 := true
											for switchFlag1 == true {
												fmt.Println("Choose one \n1)WITHDRAWAL 2)CHECK BALANCE")
												var choice int
												fmt.Scanln(&choice)
												if choice == 1 {
													fmt.Println("Enter Amount:")
													var amt int
													fmt.Scanln(&amt)
													if amt <= atmBalance {
														if amt <= dailyLimit {
															if amt <= savingAccBalance {
																fmt.Println("Collect Cash")
																fmt.Println("-------------------------------")
																savingAccBalance -= amt
																fmt.Printf("Amount withdrawn %d \nAccount Balance %d", amt, savingAccBalance)
																flag = false
																switchFlag1 = false
																forFlag = false

															} else {
																fmt.Println("Insufficent Balance")
															}
														} else {
															fmt.Println("Daily Cash Withdrawal Limit Exceeded")
														}
													} else {
														fmt.Printf("Atm Balance Insuffcient \n Only %d available", atmBalance)
													}
												} else if choice == 2 {
													fmt.Printf("Account Balance : %d", savingAccBalance)
													flag = false
													switchFlag1 = false
													forFlag = false
												} else {
													fmt.Println("Incorrect Selection Please select from given choices")
												}
											}

										case 2:
											switchFlag1 := true
											for switchFlag1 == true {
												fmt.Println("Choose one \n1)WITHDRAWAL 2)CHECK BALANCE")
												var choice int
												fmt.Scanln(&choice)
												if choice == 1 {
													fmt.Println("Enter Amount:")
													var amt int
													fmt.Scanln(&amt)
													if amt <= atmBalance {
														if amt <= dailyLimit {
															if amt <= currentAccBalance {
																fmt.Println("Collect Cash")
																fmt.Println("-------------------------------")
																currentAccBalance -= amt
																fmt.Printf("Amount withdrawn %d \nAccount Balance %d", amt, currentAccBalance)
																flag = false
																switchFlag1 = false
																forFlag = false

															} else {
																fmt.Println("Insufficent Balance")
															}
														} else {
															fmt.Println("Daily Cash Withdrawal Limit Exceeded")
														}
													} else {
														fmt.Printf("Atm Balance Insuffcient \n Only %d available", atmBalance)
													}
												} else if choice == 2 {
													fmt.Printf("Account Balance : %d", currentAccBalance)
													flag = false
													switchFlag1 = false
													forFlag = false
												} else {
													fmt.Println("Incorrect Selection Please select from given choices")
												}
											}
										default:
											fmt.Printf("Incorrect selection \n Please select only given choices")
										}

									} else {
										fmt.Println("Incorrect Selection plese select from the given choices")
										typeFlag = true
									}
								}
							} else {
								fmt.Println("Wrong Pin... Enter again")

							}
						} else {
							fmt.Println("Can enter only 4 digits pin")
							break
						}

					} else {
						fmt.Println("Invaild Card ")
						break
					}
				} else {
					fmt.Println("Couldnt read the inserted card check the card pls")
					break
				}
			} else {
				fmt.Println("Atm balance is Insufficient")
			}
		}
	}
}
