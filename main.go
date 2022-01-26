package main

import (
	"strconv"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)



func main(){
	a := app.New()  
	w := a.NewWindow("MY CALC")
	w.Resize(fyne.NewSize(300, 450))
	// w.SetFixedSize(true)
	output:=""
	appname := widget.NewLabel("Calculator By Sunidhi!!")
	input := widget.NewLabel(output)
	historyStr := ""
	history:=widget.NewLabel(historyStr)
	ishistory := false
	
	var historyARR []string; 
	historybtn:=widget.NewButton("History",func(){
		if !ishistory{
			for i:=len(historyARR)-1;i>=0;i--{
				historyStr = historyStr+historyARR[i];
				historyStr+="\n";
				ishistory = true
			}
			}else{
				historyStr = ""
				ishistory = false
			}
			history.SetText(historyStr)
		})
	backbtn:=widget.NewButton("Back",func ()  {
		if len(output) > 0{
		output = output[0:len(output)-1]	
		input.SetText(output)
	}})
	clearbtn:=widget.NewButton("Clear",func ()  {
		output =" "
		input.SetText(output)
		})
	openbtn:=widget.NewButton("(",func ()  {
		output =output+"("
		input.SetText(output)
		})
	closebtn:=widget.NewButton(")",func ()  {
		output =output+")"
		input.SetText(output)
		})
	dividebtn:=widget.NewButton("/",func ()  {
		output =output+"/"
		input.SetText(output)
		})
	sevenbutton:=widget.NewButton("7",func ()  {
		output =output+"7"
		input.SetText(output)
		})
	eightbutton:=widget.NewButton("8",func ()  {
		output =output+"8"
		input.SetText(output)
		})
	ninebutton:=widget.NewButton("9",func ()  {
		output =output+"9"
		input.SetText(output)
		})
	multiplybutton:=widget.NewButton("*",func ()  {
		output =output+"*"
		input.SetText(output)
		})
	fourbutton:=widget.NewButton("4",func ()  {
		output =output+"4"
		input.SetText(output)
		})
	fivebutton:=widget.NewButton("5",func ()  {
		output =output+"5"
		input.SetText(output)
		})
	sixbutton:=widget.NewButton("6",func ()  {
		output =output+"6"
		input.SetText(output)
		})
	minusbutton:=widget.NewButton("-",func ()  {
		output =output+"-"
		input.SetText(output)
		})
	onebutton:=widget.NewButton("1",func ()  {
		output =output+"1"
		input.SetText(output)
		})
	twobutton:=widget.NewButton("2",func ()  {
		output =output+"2"
		input.SetText(output)
		})
	threebutton:=widget.NewButton("3",func ()  {
		output =output+"3"
		input.SetText(output)
		})
	plusbutton:=widget.NewButton("+",func ()  {
		output =output+"+"
		input.SetText(output)
		})
	zerobutton:=widget.NewButton("0",func ()  {
		output =output+"0"
		input.SetText(output)
		})
	dotbutton:=widget.NewButton(".",func ()  {
		output =output+"."
		input.SetText(output)
		})
	equalbutton:=widget.NewButton("=",func ()  {
		expression, err := govaluate.NewEvaluableExpression(output);
		if err == nil{ 
			result, err := expression.Evaluate(nil);
			if err == nil{
				ans := strconv.FormatFloat(result.(float64),'f',-1,64);;
				strTobeAppend:=output+"="+ans
				historyARR = append(historyARR,strTobeAppend)
				output = ans
			}else{
				output="ERROR"
			}
		}else{
			output="ERROR"
		}
		input.SetText(output)
		output = ""
		})
	w.SetContent(container.NewVBox(
		appname,
		input,
		history,
		container.NewGridWithColumns(1,
				historybtn,backbtn,),
		container.NewGridWithColumns(4,
					clearbtn,openbtn,closebtn,dividebtn,
					sevenbutton,eightbutton,ninebutton,multiplybutton,
					fourbutton,fivebutton,sixbutton,minusbutton,
					onebutton,twobutton,threebutton,plusbutton,),
		container.NewGridWithColumns(3,
				zerobutton,dotbutton,equalbutton),
			),
		)
		w.ShowAndRun()

}