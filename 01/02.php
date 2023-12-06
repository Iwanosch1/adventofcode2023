<?php

function replaceWordNumbers($data) {
    # i know this is dirty but does the job
    $pattern1 = '/one/i';
    $pattern2 = '/two/i';
    $pattern3 = '/three/i';
    $pattern4 = '/four/i';
    $pattern5 = '/five/i';
    $pattern6 = '/six/i';
    $pattern7 = '/seven/i';
    $pattern8 = '/eight/i';
    $pattern9 = '/nine/i';

    $data = preg_replace($pattern1, 'one1one', $data);
    $data = preg_replace($pattern2, 'two2two', $data);
    $data = preg_replace($pattern3, 'three3three', $data);
    $data = preg_replace($pattern4, 'four4four', $data);
    $data = preg_replace($pattern5, 'five5five', $data);
    $data = preg_replace($pattern6, 'six6six', $data);
    $data = preg_replace($pattern7, 'seven7seven', $data);
    $data = preg_replace($pattern8, 'eight8eight', $data);
    $data = preg_replace($pattern9, 'nine9nine', $data);

    return $data;
}

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
    $input = "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen";
    $workedInput = replaceWordNumbers($input);
    getSum($workedInput);
}

// test();
$myfile = fopen("input1.txt", "r") or die("Unable to open file!");
$content = fread($myfile,filesize("input1.txt"));
$workedContent = replaceWordNumbers($content);
getSum($workedContent);
fclose($myfile);
?>