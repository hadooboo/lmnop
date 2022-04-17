import 'package:flutter/material.dart';

class UserPage extends StatelessWidget {
  const UserPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    dynamic arguments = ModalRoute.of(context)?.settings.arguments as Map;

    return Scaffold(
      body: Container(
        alignment: AlignmentDirectional.center,
        child: Text(arguments['userID'])
      ),
    );
  }
}
