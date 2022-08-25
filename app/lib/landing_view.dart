import 'package:lmnop/landing_view_model.dart';
import 'package:flutter/material.dart';
import 'package:keyboard_dismisser/keyboard_dismisser.dart';
import 'package:provider/provider.dart';

class LandingView extends StatelessWidget {
  late LandingViewModel viewModel;

  final _userIDController = TextEditingController();
  final _formKey = GlobalKey<FormState>();

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
        Form(
          key: _formKey,
          child: Padding(
            padding: const EdgeInsets.fromLTRB(40, 0, 40, 40),
            child: TextFormField(
              decoration: const InputDecoration(
                border: UnderlineInputBorder(),
                labelText: 'BOJ 아이디',
              ),
              autovalidateMode: AutovalidateMode.onUserInteraction,
              controller: _userIDController,
              validator: (text) {
                if (text == null || text.isEmpty) {
                  return '아이디를 입력해주세요.';
                }
                if (!RegExp(r'^[a-zA-Z0-9_]*$').hasMatch(text)) {
                  return '아이디 형식이 올바르지 않습니다.';
                }
                return null;
              },
            ),
          ),
        ),
        FloatingActionButton(
          child: const Icon(Icons.search),
          onPressed: () {
            if (_formKey.currentState!.validate()) {
              viewModel.setUserID(context, _userIDController.text);
            }
          },
        ),
      ],
    );
  }
}
