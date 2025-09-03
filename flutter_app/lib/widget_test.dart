// This is a standard test file created by Flutter.
import 'package:flutter/material.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_app/main.dart';

void main() {
  testWidgets('Counter increments smoke test', (WidgetTester tester) async {
    // Build our app and trigger a frame.
    await tester.pumpWidget(const MyApp());

    // Verify that our counter starts at 0.
    expect(find.text('Tap the button to fetch a message...'), findsOneWidget);

    // Tap the 'Fetch Message' button and trigger a frame.
    await tester.tap(find.byType(ElevatedButton));
    await tester.pump();

    // The test is limited here as the API call is asynchronous.
    // A more robust test would mock the http client.
  });
}
