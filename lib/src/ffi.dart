//file name ffi.dart
import 'dart:ffi' as ffi; // For FFI
import 'dart:io' show Platform;

import 'package:ffi/ffi.dart';
import 'package:ffi/src/utf8.dart';

typedef get_key_func = ffi.Pointer<Utf8> Function(); // FFI fn signature
typedef GetKey = ffi.Pointer<Utf8> Function(); // Dart fn signature

// typedef get_mtpinfo_func = ffi.Pointer<Utf8> Function(); // FFI fn signature
// typedef GetMTPInfo = ffi.Pointer<Utf8> Function(); // Dart fn signature

typedef get_file_func = ffi.Pointer<Utf8> Function(); // FFI fn signature
typedef GetFile = ffi.Pointer<Utf8> Function(); // Dart fn signature

void run() {
  // TODO: can we move this code in an init func or something? Probably have to do class BS. at least make this func a Once.
  // dyn lib file found at root of the project for now.
	ffi.DynamicLibrary dylib;
	if (Platform.isWindows) {
		dylib = ffi.DynamicLibrary.open('foo.dll');
	} else {
		dylib = ffi.DynamicLibrary.open('foo.so');
	}
	number(dylib);
	mtp(dylib);
}

void number(ffi.DynamicLibrary dylib) {
	final GetKey getKey = dylib.lookup<ffi.NativeFunction<get_key_func>>('GetKey').asFunction();

	ffi.Pointer<Utf8> theKey = getKey();
	print(theKey.toDartString());
}

void mtp(ffi.DynamicLibrary dylib) {
	// final GetMTPInfo getMTPInfo = dylib.lookup<ffi.NativeFunction<get_mtpinfo_func>>('GetMTPInfo').asFunction();
	final GetFile getFile = dylib.lookup<ffi.NativeFunction<get_file_func>>('GetFile').asFunction();

	ffi.Pointer<Utf8> theFile = getFile();
	print(theFile.toDartString());

//	ffi.Pointer<Utf8> mtpInfo = getMTPInfo();
//	print(mtpInfo.toDartString());
}
