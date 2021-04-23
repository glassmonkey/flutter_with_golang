import 'package:pigeon/pigeon.dart';

class Request {
  late String expression;
}

// 戻り値の定義
class Response {
  late String message;
}

@HostApi()
abstract class Api {
  Response call();
  Response eval(Request req);
}

@FlutterApi()
abstract class CallbackApi {
  void apply(Response response);
}
