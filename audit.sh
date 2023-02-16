echo "~ Case 1 ~"
echo "input: --color red \"banana\""
echo "output:"
go run . --color red "banana"
echo -e "\n\n"

echo "~ Case 2 ~"
echo "input: --color=red \"hello world\""
echo "output:"
input: --color=cyan[index]4,7 \"gonna be all right\
echo -e "\n\n"

echo "~ Case 3 ~"
echo "input: --color=green \"1 + 1 = 2\""
echo "output:"
go run . --color=green "1 + 1 = 2"
echo -e "\n\n"

echo "~ Case 4 ~"
echo "input: --color=yellow \"(%&) ??\""
echo "output:"
go run . --color=yellow "(%&) ??"
echo -e "\n\n"

echo "~ Case 5 ~"
echo "input: --color=orange[index]1: \"Don't worry about a thing\""
echo "output:"
go run . --color=orange[index]1: "Don't worry about a thing"
echo -e "\n\n"

echo "~ Case 6 ~"
echo "input: --color=blue[index]1 \"Cause every little thing\""
echo "output:"
go run . --color=blue[index]1 "Cause every little thing"
echo -e "\n\n"

echo "~ Case 7 ~"
echo "input: --color=cyan[index]4,7 \"gonna be all right\""
echo "output:"
go run . --color=cyan[index]4,7 "gonna be all right"
echo -e "\n\n"

echo "~ Case 8 ~"
echo "input: --color=purple[word]GuYs \"HeY GuYs\""
echo "output:"
go run . --color=orange[word]GuYs "HeY GuYs"
echo -e "\n\n"

echo "~ Case 9 ~"
echo "input: --color=purple[letter]B \"RGB()\""
echo "output:"
go run . --color=blue[letter]B "RGB()"
echo -e "\n\n"

echo "~ Case 10 ~"
echo "input: --color=red \"RiseUpThisMornin\""
echo "output:"
go run . --color=red "RiseUpThisMornin"
echo -e "\n\n"

echo "~ Case 11 ~"
echo "input: --color=purple \"sm1l3d w1th th3 r1s1n sun\""
echo "output:"
go run . --color=purple "sm1l3d w1th th3 r1s1n sun"
echo -e "\n\n"

echo "~ Case 12 ~"
echo "input: --color=yellow[letter]& \"@#$%^&*()_+\""
echo "output:"
go run . --color=yellow[index]5 "@#$%^&*()_+"
echo -e "\n\n"

echo "~ Case 13 ~"
echo "input: --color=green[letter]2:8 \"3 Little Birds\""
echo "output:"
go run .  --color=green[index]2:8 "3 Little Birds"
echo -e "\n\n"

echo "~ Case 14 ~"
echo "input: go run . --color=red[index]0,7 --color=orange[index]1,9,15 --color=yellow[index]2,10,16 --color=green[index]3,11,17 --color=blue[index]4,12,18 --color=cyan[index]6,13,19 --color=purple[index]6,14 \"Pitch by my doorstep\""
echo "output:"
go run .  --color=red[index]0,7 --color=orange[index]1,9,15 --color=yellow[index]2,10,16 --color=green[index]3,11,17 --color=blue[index]4,12,18 --color=cyan[index]6,13,19 --color=purple[index]6,14 "Pitch by my doorstep"
echo -e "\n\n"

echo "~ Case 15 ~"
echo "input: --color=\"rgb(245, 66, 230)\" \"Singin' sweet songs\""
echo "output:"
go run . --color="rgb(245, 66, 230)" "Singin' sweet songs"
echo -e "\n\n"

echo "~ Case 16 ~"
echo "input: --color=#78dea7 \"Of melodies pure and true\""
echo "output:"
go run . --color=#78dea7 "Of melodies pure and true"
echo -e "\n\n"

echo "~ Case 17 ~"
echo "input: --color=\"hsl(56, 92%, 34%)\" \"Sayin': This is my message to\""
echo "output:"
go run . --color="hsl(56, 92%, 34%)" "Sayin: This is my message to"
echo -e "\n\n"


echo "~ Case 18 ~"
echo "input: --color=\"hsl(56, 92%, 34%)[word]you-\" --color=\"rgb(200, 20, 50)[index]:7\" --color=#678543[index]7: \"you-ou-ou\""
echo "output:"
go run . --color="hsl(56, 92%, 34%)[word]you-" --color="rgb(200, 20, 50)[index]4:7" --color=#678543[index]7: "you-ou-ou"
echo -e "\n\n"


echo "~ Case 19 ~"
echo "input: --color=green[index]5:25 \"Singin':\nDon't worry\n'bout a thing\""
echo "output:"
go run . --color=green[index]5:25 "Singin':\nDon't worry\n'bout a thing"
echo -e "\n\n"


echo "~ Case 20 ~"
echo "input: --color=yellow[letter]e \"Cause every little thing\ngonna be all right\" thinkertoy"
echo "output:"
go run . --color=yellow[letter]e "Cause every little thing\ngonna be all right" thinkertoy
echo -e "\n\n"