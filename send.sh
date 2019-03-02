
#!/bin/sh

allThreads=(
"Sigma Alpha"
"Sigma Alpha Iota"
"Sigma Gamma Tau"
"Sigma Lambda Chi"
"Sigma Phi Epsilon"
"Sigma Pi"
"Society of Women in Physics and Astronomy"
"Spanish Club"
"Student Government"
"Student Union Board"
"Student Volunteer Services"
"The Pride Alliance"
"The Yoga Club"
"Theta Chi"
"UNICEF"
"User Experience Design Club"
"Web Development Club"
"Women in Aviation"
"Women Who Design"
"Zeta Phi Beta Sorority "
)

count=0

dir=$1

if [ -n "$1" ]; then
    echo "Dir is $1"
else
    echo "Dir cannot be empty"
    exit 1
fi

echo "#!/bin/sh" >> ./test.sh
echo "" >> ./test.sh

for i in "${allThreads[@]}"; do
    echo $count
    echo $i
    echo "./main -f incomingdata/${dir}/dataminer\ \($count\).csv -A \"$i\"; sleep 3;" >> ./test.sh
    let count=count+1
done

