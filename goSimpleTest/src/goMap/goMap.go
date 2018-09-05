package main

func main() {

}

// golang判断key是否在map中, 另外golang也没有提供item是否在array当中的判断方法,如果程序里面频繁用到了这种判断,可以将array转化为以array当中的成员为key的map再用上面的方法进行判断,这样会提高判断的效率.
func main1() {
	if _, ok := map[key]; ok {
		//存在
	}
}