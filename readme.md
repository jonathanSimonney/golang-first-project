Ex A was used to make a calculator, <br>
ex B to give some stats on the letters composing a file, <br>
and ex C to print all email adresses from all URL given in a formatted file (only url, 
with one url per line).
To use, install golang interpreter, and use : <br>
 -go run exA.go operand number number2 ... numberN (print result of operation ; if you want to multiply, do not forget to write "*", not simply *).<br><br>
 -go run exB.go fileToAnalyse (print letters with most and least occurences, and respective number of occurence)<br><br>
 -go run exC.go fileWithListUrl (print number of @ in url listed, and display number and list of email adresses in url in file given.)
 exC.go can be called with third argument. If this third argument is y (the letter), duplicate email adresses will be counted only once.