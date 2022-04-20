import 'package:flutter/material.dart';
import 'package:lmnop/user_view.dart';
import 'package:lmnop/user_view_model.dart';
import 'package:provider/provider.dart';

class UserPage extends StatelessWidget {
  const UserPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    dynamic arguments = ModalRoute.of(context)?.settings.arguments as Map;

    return ChangeNotifierProvider(
      create: (_) => UserViewModel(arguments['userID']),
      child: UserView(),
    );
  }
}
