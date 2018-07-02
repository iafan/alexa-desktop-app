#!/bin/sh

ROOT=$HOME/sdk-folder

cd $ROOT/sdk-build
cmake \
    $ROOT/sdk-source/avs-device-sdk \
    -DCMAKE_BUILD_TYPE=DEBUG \
    -DPORTAUDIO=ON \
    -DPORTAUDIO_LIB_PATH=$ROOT/third-party/portaudio/lib/.libs/libportaudio.a \
    -DPORTAUDIO_INCLUDE_DIR=$ROOT/third-party/portaudio/include \
    -DKITTAI_KEY_WORD_DETECTOR=ON \
    -DKITTAI_KEY_WORD_DETECTOR_LIB_PATH=$ROOT/third-party/snowboy/lib/osx/libsnowboy-detect.a \
    -DKITTAI_KEY_WORD_DETECTOR_INCLUDE_DIR=$ROOT/third-party/snowboy/include

make SampleApp -j2