import 'package:flutter/material.dart';
import 'package:lmnop/user_view_model.dart';
import 'package:provider/provider.dart';

class UserView extends StatelessWidget {
  late UserViewModel viewModel;

  @override
  Widget build(BuildContext context) {
    viewModel = Provider.of<UserViewModel>(context);
    return Scaffold(
      body: _buildUserViewBody(context),
    );
  }

  Widget _buildUserViewBody(BuildContext context) {
    return Padding(
      padding: EdgeInsets.fromLTRB(20, 20 + MediaQuery.of(context).padding.top, 20, 20),
      child: Column(
        mainAxisAlignment: MainAxisAlignment.start,
        crossAxisAlignment: CrossAxisAlignment.start,
        children: [
          Text(viewModel.user?.handle ?? ""),
          Text("총 문제 해결: ${viewModel.user?.solvedCount}문제"),
          Text("최대 연속 문제 해결: ${viewModel.user?.maxStreak}일"),
          Text("${viewModel.user?.tier} ${viewModel.user?.rating}"),
          Text("${viewModel.user?.classNumber} ${viewModel.user?.classDecoration}"),
          const Text("추천 문제"),
          Text("문제 번호: ${viewModel.problem?.problemId}"),
          Text("문제 레벨: ${viewModel.problem?.level}"),
        ],
      ),
    );
  }
}
