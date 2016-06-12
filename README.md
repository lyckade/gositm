# gositm
A very simple time machine written in golang

## How it works?

gositm walks throug the given folder tree and adds to each file a timestamp of the last change date and time. Then it copies the file with the timestampt inside the name into a backup folder.

Inside the backup folder all versions of the file are still exist. The whole folder structure is also availialbe inside the backup folder. To browse all the file versions there is no special tool needed.
