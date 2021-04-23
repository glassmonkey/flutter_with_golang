import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:flutter_with_golang/view_models/message.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';

class EvalWidget extends HookWidget {
  @override
  Widget build(BuildContext context) {
    final messageState = useProvider(messageProvider);
    final messageNotifier = useProvider(messageProvider.notifier);
    final textController = TextEditingController();

    return Scaffold(
      appBar: AppBar(
        title: const Text('Sample Text Calculator'),
      ),
      body: Center(
        child: Column(
          mainAxisAlignment: MainAxisAlignment.center,
          children: <Widget>[
            TextField(
              controller: textController,
            ),
            Text('result: ${messageState.text}'),
          ],
        ),
      ),
      floatingActionButton: FloatingActionButton(
        onPressed: () {
          messageNotifier.evaluation(textController.text);
        },
        tooltip: 'Culculate',
        child: Icon(Icons.transit_enterexit_sharp),
      ),
    );
  }
}
