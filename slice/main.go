package main

func main()  {
	nums := make([]int, 10)

	for i := 0; i < 10;i ++{
		index := i
		go func() {
			nums[index] = index
		}()
	}

	println(nums)
}
