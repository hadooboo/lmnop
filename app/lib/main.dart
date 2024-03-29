import 'package:flutter/services.dart';
import 'package:lmnop/routes.dart';
import 'package:flutter/material.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatelessWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    SystemChrome.setPreferredOrientations([
      DeviceOrientation.portraitUp,
      DeviceOrientation.portraitDown,
    ]);
    return MaterialApp(
      title: 'My Optimum Problem',
      theme: ThemeData(
        primarySwatch: Colors.green,
      ),
      routes: routes,
    );
  }
}
