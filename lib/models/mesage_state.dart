class MessageState {
  final String text;
  MessageState({required this.text});

  MessageState copyWith({required String text}) {
    return MessageState(text: text);
  }
}
