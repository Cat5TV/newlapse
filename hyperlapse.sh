#!/bin/bash
clear
echo ""
echo 'Recording hyperlapse...'
echo 'When finished, press CTRL-C to save video'
echo 'But make sure you ONLY PRESS IT ONCE'

folder='/tmp/hyperlapse' # Set to what you'd like... /home/robbie/hyperlapose for example
cd $folder

newlapse -ccc -rate 1 -fps 24 -folder "$folder" > /dev/null 2>&1

tar -czvf $folder/backup-`date +"%Y-%m-%d-%T"`.tar.gz $folder/*.jpg $folder/1s
rm $folder/*.jpg
rm -rf $folder/1s

name=hyperlapse-`date +%Y-%m-%d-%T`
if [[ -e $name.mp4 ]] ; then
    i=0
    while [[ -e $name-$i.mp4 ]] ; do
        let i++
    done
    name=$name-$i
fi
mv video_1.mp4 "$name".mp4

echo ""
echo DONE
echo File is saved in folder "hyperlapse"
sleep 5
