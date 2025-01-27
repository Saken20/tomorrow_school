interview=$(head -n "179" "streets/Buckingham_Place" | tail -n -1)
read -a array <<< "$interview"
export INTERVIEW_NUMBER=$(echo "${array[2]#?}")

interview=$(cat "interviews/interview-${INTERVIEW_NUMBER}")

echo "$INTERVIEW_NUMBER"
echo "$interview"
echo "$MAIN_SUSPECT"