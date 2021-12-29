package main

type Item struct {
	iValue int
	pNext  *Item
}

func main() {

}

func getItem(pHead *Item, m int) (iRet *Item, err error) {
	//时间复杂度低，空间复杂度高
	pH := pHead
	pTmp := pH
	plog := pHead
	itemLen := 0
	storeList := make(map[int]*Item)
	if 0 == m {
		return pHead, nil
	}
	for {
		if nil != pTmp {

			//如果itemLen>m,plog ++ ,结尾

			storeList[itemLen] = pTmp
			itemLen++
			pTmp = pTmp.pNext
		} else {
			break
		}
	}

}

//func getItem(pHead *Item, m int) (iRet *Item,err error)  {
//时间复杂度高，空间复杂度低
//pH := pHead
//pTmp := pH
//itemLen := 0
//if 0 == m {
//	return pHead,nil
//}
//for{
//	if nil != pTmp{
//		itemLen ++
//		pTmp = pTmp.pNext
//	}else {
//		break
//	}
//}
//if m > itemLen{
//	return nil, err
//}
//pTmp = pHead
//num := itemLen - m
//for i:=0;i<num;i++{
//	pTmp = pTmp.pNext
//}
//return pTmp,nil

//}
