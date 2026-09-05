#!/bin/bash
# Use ./new_dir.sh CHAPTER_DIR_NAME NEW_LESSON_DIR_NAME
CHAPTER_DIR=${1}
LESSON_DIR=${2}

SCRIPT_DIR=$(cd "$(dirname "$0")" && pwd)
cd ${SCRIPT_DIR}

cd ${CHAPTER_DIR} || { echo "Making it now!" && mkdir ${CHAPTER_DIR} && cd ${CHAPTER_DIR}; }
cp -r ../template ${LESSON_DIR}/
cd ${LESSON_DIR}
mv temp.go ${LESSON_DIR}.go
mv temp_test.go ${LESSON_DIR}_test.go

