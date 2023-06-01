//file name fficheck.dart
import 'dart:ffi' as ffi; // For FFI
import 'package:ffi/ffi.dart';
import 'package:ffi/src/utf8.dart';

// TODO: on macos, foo.so is at the root of the project for now
// TODO: conditional code for foo.so VS foo.dll
// final dylib = ffi.DynamicLibrary.open('foo.so');
final dylib = ffi.DynamicLibrary.open('foo.dll');

typedef get_key_func = ffi.Pointer<Utf8> Function(); // FFI fn signature
typedef GetKey = ffi.Pointer<Utf8> Function(); // Dart fn signature

// typedef get_mtpinfo_func = ffi.Pointer<Utf8> Function(); // FFI fn signature
// typedef GetMTPInfo = ffi.Pointer<Utf8> Function(); // Dart fn signature

typedef get_file_func = ffi.Pointer<Utf8> Function(); // FFI fn signature
typedef GetFile = ffi.Pointer<Utf8> Function(); // Dart fn signature

final GetKey getKey = dylib.lookup<ffi.NativeFunction<get_key_func>>('GetKey').asFunction();
// final GetMTPInfo getMTPInfo = dylib.lookup<ffi.NativeFunction<get_mtpinfo_func>>('GetMTPInfo').asFunction();
final GetFile getFile = dylib.lookup<ffi.NativeFunction<get_file_func>>('GetFile').asFunction();

void testffi() {
	print("TESTFFI");
	ffi.Pointer<Utf8> theKey = getKey();
	print(theKey.toDartString());

	ffi.Pointer<Utf8> theFile = getFile();
	print(theFile.toDartString());

//	ffi.Pointer<Utf8> mtpInfo = getMTPInfo();
//	print(mtpInfo.toDartString());
}
