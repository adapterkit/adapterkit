string=$(find examples | grep go.mod | sed 's/go.mod//')

list=()
for i in $string; do
    list+=("$i")
done

len=${#list[@]}
printf "::set-output name=matrix::{\"include\": [\n"
for ((i = 0; i < len; i++)); do
    if [ "$i" -eq $((len - 1)) ]; then
        printf "\"%s\"\n" "${list[$i]}"
    else
        printf "\"%s\",\n" "${list[$i]}"
    fi
done

printf "]}\n"
