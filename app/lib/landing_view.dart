import 'package:lmnop/landing_view_model.dart';
import 'package:flutter/material.dart';
import 'package:keyboard_dismisser/keyboard_dismisser.dart';
import 'package:provider/provider.dart';

class LandingView extends StatelessWidget {
  late LandingViewModel viewModel;

  final userIDController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    viewModel = Provider.of<LandingViewModel>(context);
    return KeyboardDismisser(
      gestures: const [GestureType.onTap],
      child: Scaffold(
        body: _buildLandingViewBody(context),
      ),
    );
  }

  Widget _buildLandingViewBody(BuildContext context) {
    return Column(
      mainAxisAlignment: MainAxisAlignment.center,
      children: [
        const Image(image: AssetImage('images/lmnop_logo.png')),
        Padding(
          padding: const EdgeInsets.fromLTRB(40, 0, 40, 40),
          child: TextField(
            decoration: const InputDecoration(
              border: UnderlineInputBorder(),
              labelText: 'BOJ 아이디',
            ),
            controller: userIDController,
          ),
        ),
        FloatingActionButton(
          child: const Icon(Icons.search),
          onPressed: () => viewModel.setUserID(context, userIDController.text),
        ),
      ],
    );
  }
}
