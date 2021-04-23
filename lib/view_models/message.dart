import 'package:flutter_with_golang/models/mesage_state.dart';
import 'package:flutter_with_golang/pigeon/native.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

final messageProvider = StateNotifierProvider<MessageNotifier, MessageState>(
    (_) => MessageNotifier(client: Api()));

class MessageNotifier extends StateNotifier<MessageState> {
  final Api client;

  MessageNotifier({required this.client}) : super(MessageState(text: "result"));

  call() async {
    final response = await this.client.call();
    this.state = this.state.copyWith(text: response.message ?? "result");
  }

  evaluation(String text) async {
    final req = Request();
    req.expression = text;
    final response = await this.client.eval(req);
    this.state = this.state.copyWith(text: response.message ?? "result");
  }
}
