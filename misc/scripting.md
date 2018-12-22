scripting
---

[valid-phone-number](https://www.interviewbit.com/problems/valid-phone-number/)

```bash

#!/bin/bash

regex1='[0-9]{3}-[0-9]{3}-[0-9]{4}$'
regex2='\([0-9]{3}\)\s[0-9]{3}-[0-9]{4}$'

cat input | while read LINE
do
    if [[ $LINE =~ $regex1|$regex2 ]];
    then
	echo $LINE
    fi
done
```

[Lines in a given range](https://www.interviewbit.com/problems/lines-in-a-given-range/)


```bash
#!/bin/bash

RANGE=$(head -n 1 'input')
L=$(awk {'print $1'} <<< $RANGE)
R=$(awk {'print $2'} <<< $RANGE)
cat input | head -n $R | tail -n $((R-L+1))
```
