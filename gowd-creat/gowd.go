package main

import (
	"fmt"
	"log"
	"os"
)

// lists
var (
	onlynumbers              = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	onlylowercase            = []string{"q", "w", "e", "r", "t", "y", "u", "i", "o", "p", "a", "s", "d", "f", "g", "h", "j", "k", "l", "z", "x", "c", "v", "b", "n", "m"}
	onlyuppercase            = []string{"Q", "W", "E", "R", "T", "Y", "U", "I", "O", "P", "A", "S", "D", "F", "G", "H", "J", "K", "L", "Z", "X", "C", "V", "B", "N", "M"}
	onlyLetters              = append(onlylowercase, onlyuppercase...)
	smallspecialnotlisted    = []string{"!", "?", "*", "/", "+"}
	charsetwithnumbers       = append(onlyLetters, onlynumbers...)
	speicalcharset           = []string{"!", "?", "*", "/", "-", "_", "^", "<", ">", "$", "+", "-", "&", "@"}
	charsetWithspecial       = append(onlyLetters, speicalcharset...)
	biggestcharset           = append(charsetWithspecial, onlynumbers...)
	charsetwithsmallspecials = append(smallspecialnotlisted, onlyLetters...)
)

// other var's
var outputlines = []string{}

func printlist() {
	fmt.Println("only numbers: ", onlynumbers)
	fmt.Println("only downcase: ", onlylowercase)
	fmt.Println("only uppercase: ", onlyuppercase)
	fmt.Println("only letters: ", onlyLetters)
	fmt.Println("numbers and letters: ", charsetwithnumbers)
	fmt.Println("special charset: ", speicalcharset)
	fmt.Println("letters with special charset: ", charsetWithspecial)
	fmt.Println("biggest charset: ", biggestcharset)
	fmt.Println("less special character with all list", charsetwithsmallspecials)
}

func userplace() {
	const version float32 = 2.3
	multiLineString := `
  ////////////////by k3sR4T////////////////
          GOWd-wordliss-creator
  create wordlist for brute forcing attacks
  creat wordlist for all posibilities`
	fmt.Println(multiLineString)
	fmt.Println("    ////////////", "version:", float32(version), "////////////")

	options := `
	1 use only numbers
	2 use only lowercase 
	3 use only uppercase
	4 use all letters
	5 use numbers and letters 
	6 use special charset
	7 use onlyletters with special charset
	8 use biggest charset
	9 charset with small special
	0 print list && exit
///////////////////////////////////////	
	`
	fmt.Println(options)
}

func main() {
	userplace()
	var userInput int8 = 99
	fmt.Print("sellect an option: ")
	fmt.Scanln(&userInput)
	fmt.Println("entered:", userInput)
	var resultList []string
	if userInput == 0 {
		printlist()
		os.Exit(0)
	} else if userInput == 1 {
		resultList = onlynumbers
	} else if userInput == 2 {
		resultList = onlylowercase
	} else if userInput == 3 {
		resultList = onlyuppercase
	} else if userInput == 4 {
		resultList = onlyLetters
	} else if userInput == 5 {
		resultList = charsetwithnumbers
	} else if userInput == 6 {
		resultList = speicalcharset
	} else if userInput == 7 {
		resultList = charsetWithspecial
	} else if userInput == 8 {
		resultList = biggestcharset
	} else if userInput == 9 {
		resultList = charsetwithsmallspecials
	} else if userInput >= 10 {
		fmt.Println("please select defined list number")
		fmt.Println("quiting")
		os.Exit(0)
	} else {
		fmt.Println("please select defined mod")
		fmt.Println("quiting")
		os.Exit(0)
	}
	var list = []string(resultList)

	var userInputTwo int8 = 1
	fmt.Print("how long is your password( 1-12): ")
	fmt.Scanln(&userInputTwo)
	fmt.Println("long: ", userInputTwo)

	if userInputTwo == 0 || userInputTwo == 1 {
		fmt.Println(list)

	} else if userInputTwo == 2 {
		stopper_value := list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				fmt.Println(item1 + item2)
				outputlines = append(outputlines, item1+item2)
				if item1+item2 == stopper_value {
					break
				}
			}
		}

	} else if userInputTwo == 3 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					fmt.Println(item1 + item2 + item3)
					outputlines = append(outputlines, item1+item2+item3)
					if item1+item2+item3 == stopper_value {
						break
					}
				}
			}
		}

	} else if userInputTwo == 4 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						fmt.Println(item1 + item2 + item3 + item4)
						outputlines = append(outputlines, item1+item2+item3+item4)
						if item1+item2+item3+item4 == stopper_value {
							break
						}

					}
				}
			}
		}

	} else if userInputTwo == 5 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						for _, item5 := range list {
							fmt.Println(item1 + item2 + item3 + item4 + item5)
							outputlines = append(outputlines, item1+item2+item3+item4+item5)
							if item1+item2+item3+item4+item5 == stopper_value {
								break
							}
						}
					}
				}
			}
		}

	} else if userInputTwo == 6 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						for _, item5 := range list {
							for _, item6 := range list {
								fmt.Println(item1 + item2 + item3 + item4 + item5 + item6)
								outputlines = append(outputlines, item1+item2+item3+item4+item5+item6)
								if item1+item2+item3+item4+item5+item6 == stopper_value {
									break
								}
							}
						}
					}
				}
			}
		}

	} else if userInputTwo == 7 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						for _, item5 := range list {
							for _, item6 := range list {
								for _, item7 := range list {
									fmt.Println(item1 + item2 + item3 + item4 + item5 + item6 + item7)
									outputlines = append(outputlines, item1+item2+item3+item4+item5+item6+item7)
									if item1+item2+item3+item4+item5+item6+item7 == stopper_value {
										break
									}
								}
							}
						}
					}
				}
			}
		}

	} else if userInputTwo == 8 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						for _, item5 := range list {
							for _, item6 := range list {
								for _, item7 := range list {
									for _, item8 := range list {
										fmt.Println(item1 + item2 + item3 + item4 + item5 + item6 + item7 + item8)
										outputlines = append(outputlines, item1+item2+item3+item4+item5+item6+item7+item8)
										if item1+item2+item3+item4+item5+item6+item7+item8 == stopper_value {
											break
										}
									}
								}
							}
						}
					}
				}
			}
		}

	} else if userInputTwo == 9 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						for _, item5 := range list {
							for _, item6 := range list {
								for _, item7 := range list {
									for _, item8 := range list {
										for _, item9 := range list {
											fmt.Println(item1 + item2 + item3 + item4 + item5 + item6 + item7 + item8 + item9)
											outputlines = append(outputlines, item1+item2+item3+item4+item5+item6+item7+item8+item9)
											if item1+item2+item3+item4+item5+item6+item7+item8+item9 == stopper_value {
												break
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

	} else if userInputTwo == 10 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						for _, item5 := range list {
							for _, item6 := range list {
								for _, item7 := range list {
									for _, item8 := range list {
										for _, item9 := range list {
											for _, item10 := range list {
												fmt.Println(item1 + item2 + item3 + item4 + item5 + item6 + item7 + item8 + item9 + item10)
												outputlines = append(outputlines, item1+item2+item3+item4+item5+item6+item7+item8+item9+item10)
												if item1+item2+item3+item4+item5+item6+item7+item8+item9+item10 == stopper_value {
													break
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

	} else if userInputTwo == 11 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						for _, item5 := range list {
							for _, item6 := range list {
								for _, item7 := range list {
									for _, item8 := range list {
										for _, item9 := range list {
											for _, item10 := range list {
												for _, item11 := range list {
													fmt.Println(item1 + item2 + item3 + item4 + item5 + item6 + item7 + item8 + item9 + item10 + item11)
													outputlines = append(outputlines, item1+item2+item3+item4+item5+item6+item7+item8+item9+item10+item11)
													if item1+item2+item3+item4+item5+item6+item7+item8+item9+item10+item11 == stopper_value {
														break
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

	} else if userInputTwo == 12 {
		stopper_value := list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1] + list[len(list)-1]
		for _, item1 := range list {
			for _, item2 := range list {
				for _, item3 := range list {
					for _, item4 := range list {
						for _, item5 := range list {
							for _, item6 := range list {
								for _, item7 := range list {
									for _, item8 := range list {
										for _, item9 := range list {
											for _, item10 := range list {
												for _, item11 := range list {
													for _, item12 := range list {
														fmt.Println(item1 + item2 + item3 + item4 + item5 + item6 + item7 + item8 + item9 + item10 + item11 + item12)
														outputlines = append(outputlines, item1+item2+item3+item4+item5+item6+item7+item8+item9+item10+item11+item12)
														if item1+item2+item3+item4+item5+item6+item7+item8+item9+item10+item11+item12 == stopper_value {
															break
														}
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

	} else {
		fmt.Println("not supported")
	}

	f, err := os.Create("wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for _, outputline := range outputlines {
		_, err := f.WriteString(outputline + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

	f, err = os.Create("wordlist.txt")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file
	defer f.Close()

	for _, outputline := range outputlines {
		_, err := f.WriteString(outputline + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}

}
