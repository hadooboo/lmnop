import 'package:lmnop/landing_view.dart';
import 'package:lmnop/landing_view_model.dart';
import 'package:flutter/material.dart';
import 'package:provider/provider.dart';

class LandingPage extends StatelessWidget {
  const LandingPage({Key? key}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return ChangeNotifierProvider(
      create: (_) => LandingViewModel(),
      child: LandingView(),
    );
  }
}
