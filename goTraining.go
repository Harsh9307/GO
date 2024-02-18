package main

import "fmt"

func main() {
	// var a uint8
	// a = 10
	// var student1 string = "John" // type is string
	// var student2 = "Jane"        // type is inferred
	// x := 2                       // type is inferred

	// fmt.Println(student1)
	// fmt.Println(student2)
	// fmt.Println(x)
	// fmt.Println(a)

	// fmt.Print(a, student1, student2)
	// fmt.Printf("\n")
	// fmt.Println(a, student1, student2)

	// var x1 float64 = 20.0
	// y := 42
	// fmt.Printf("x = %f\n", x1)
	// fmt.Printf("y = %d\n", y)
	// fmt.Printf("x is of type %T\n", x)
	// fmt.Printf("y is of type %T\n", y)

	// var a1 string
	// var b int
	// var c bool

	// c = true
	// fmt.Println(a1)
	// fmt.Println(b)
	// fmt.Println(c)

	// var d, e, f = 3, 4, "foo"
	// fmt.Println(d)
	// fmt.Println(e)
	// fmt.Println(f)

	// var a2, b2, c2, d2 = 1, 3, 5, 7
	// fmt.Println(a2)
	// fmt.Println(b2)
	// fmt.Println(c2)
	// fmt.Println(d2)

	// var a3, b3 = 6, "Hello"
	// c3, d3 := 7, "World"

	// fmt.Println(a3)
	// fmt.Println(b3)
	// fmt.Println(c3)
	// fmt.Println(d3)

	// var (
	// 	a int
	// 	b int    = 1
	// 	c string = "Hello"
	// )
	// fmt.Println(a)
	// fmt.Println(b)
	// fmt.Println(c)

	// var initialName, finalName string
	// fmt.Printf("Enter your first name:")
	// fmt.Scanf("%s %s", &initialName, &finalName)
	// // fmt.Printf("Enter your surname:")
	// // fmt.Scanf("%s", &finalName)
	// fmt.Printf("Your full name is %s %s\n", initialName, finalName)
	// fmt.Print("Your full name  is", initialName, " ", finalName, "\n")
	// fmt.Println("Your full name is", initialName, " ", finalName)

	// var a,b int =10,20
	// fmt.Println(a,b)
	// fmt.Println(a+b)

	// var name, age, branch, clgName string
	// fmt.Printf("Enter your name:")
	// fmt.Scanln(&name)
	// fmt.Printf("Enter your age:")
	// fmt.Scanln(&age)
	// fmt.Printf("Enter your branch:")
	// fmt.Scanln(&branch)
	// fmt.Printf("Enter your college Name:")
	// fmt.Scanln(&clgName)

	// fmt.Println("***********Printting Details***********")
	// fmt.Printf("Name : %s\n", name)
	// fmt.Printf("Age : %s\n", age)
	// fmt.Printf("Branch  : %s\n", branch)
	// fmt.Printf("College Name : %s\n", clgName)

	// fmt.Println("***********Printting Details in another way***********")
	// fmt.Printf("Name : %s   Age : %s\n", name, age)
	// fmt.Printf("Branch : %s College Name : %s\n", branch, clgName)

	// var a string = "Harsh"
	// fmt.Printf("value of a is :'%s'\n", a)

	// var str = "Geekforgeeks"
	// fmt.Printf("The string is %s \n", str)
	// var num1 int = 21
	// fmt.Printf("The decimal value is %d \n", num1)
	// var num2 float32 = 7.786
	// fmt.Printf("The floating point  value is %g %.3f\n", num2, num2)
	// var num3 int = 14
	// fmt.Printf("The binary value is %b \n", num3)
	// var num4 = 4 + 1i
	// fmt.Printf("Scientific Notation of num4 %e \n", num4)

	println(15 == 017)
	println(15 == 0xF)

	//There are four integer literal forms, the deimal(base 10 ) form , the octal (base 8) form , the hex (base 16) form and the binary form(base 2). For example, the following four integer lierals all denote

	//15 in decimal .0xF  // the hex form(starts with a "0x" or "0X")
	/*
		the octal form ( starts with a "0" ,"0o" or "0O")
		0o17 0O17 0b1111 // the binary form (starts with a "0b" or "0B")
		0B1111 15 // the decimal form (starts with a "0")
	*/

	var a int = 37
	a++
	println(a)

	fmt.Printf("The binary of %d is %b\n", a, a)
	fmt.Printf("The hexadecimal of %d is %0x\n", a, a)
	fmt.Printf("The octal of %d is %0o\n", a, a)

	/*
			Format Specifiers:

		It tells Golang how to format different data types. Some of the most commonly used specifiers are:

		⚫v-formats the value in a default format

		⚫d-formats decimal integers

		⚫g-formats the floating-point numbers

		b-formats base 2 numbers

		⚫o-formats base 8 numbers

		⚫t-formats true or false values(Boolean)
	*/
	fmt.Printf("The expression 15==oxf is %t\n", 15 == 0xf)
	fmt.Printf("The expression 15==oxf is %T\n", 15 == 0xf)

	// var i, j string = "Hello", "world"
	// fmt.Print(i, "\n")
	// fmt.Print(j, "\n")

	// var i, j = 10, "world"
	// fmt.Print(i, j)

	// var i = 15.5
	// var text = "Hello world!"

	// fmt.Printf("%v\n", i)
	// fmt.Printf("%#v\n", i)
	// fmt.Printf("%v%%\n", i)
	// fmt.Printf("%T\n", i)

	// fmt.Printf("%v\n", text)
	// fmt.Printf("%#v\n", text)
	// fmt.Printf("%v%%\n", text)
	// fmt.Printf("%T\n", text)

	var i = 15
	fmt.Printf("%b\n", i)
	fmt.Printf("%d\n", i)
	fmt.Printf("%+d\n", i)
	fmt.Printf("%o\n", i)
	fmt.Printf("%O\n", i)
	fmt.Printf("%x\n", i)
	fmt.Printf("%X\n", i)
	fmt.Printf("%#x\n", i)
	fmt.Printf("%4d\n", i)
	fmt.Printf("%-4d\n", i)
	fmt.Printf("%04d\n", i)
}
