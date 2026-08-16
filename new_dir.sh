#!/bin/bash
# Use ./new_dir.sh EXISTING_CHAPTER_DIR_NAME NEW_LESSON_DIR_NAME
CHAPTER_DIR=${1}
LESSON_DIR=${2}

cd /c/Users/Anthony/Desktop/Code/boot.dev/learn-go
cd ${CHAPTER_DIR}
cp -r ../template ${LESSON_DIR}/
cd ${LESSON_DIR}
mv temp.go ${LESSON_DIR}.go
mv temp_test.go ${LESSON_DIR}_test.go

