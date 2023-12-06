<?php

function getSum($data) {
    $separator = "\r\n";
    $line = strtok($data, $separator);

    $sum = 0;
    while ($line !== false) {
        preg_match_all('/[0-9]/', $line, $numbers); # get each line
        $numbers = array_filter($numbers); # get numbers in array

        $firstNum = reset($numbers[0]); # get first number
        $lastNum = end($numbers[0]); # get last number

        $lineNumber = $firstNum * 10 + $lastNum; # combine numbers        
        $sum = $sum + $lineNumber;
        # get next line
        $line = strtok( $separator );
    }
    echo "final sum: " . $sum;
}
function test() {
    $input = "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet\nasefe1\nasef11pogh";

    getSum($input);
}

// test();
$myfile = fopen("input1.txt", "r") or die("Unable to open file!");
getSum(fread($myfile,filesize("input1.txt")));
fclose($myfile);

# 12142132 is too high
?>