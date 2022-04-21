import 'package:flutter/material.dart';
import 'package:lmnop/user_view_model.dart';
import 'package:lmnop/constants.dart';
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
    return Container(
      width: double.infinity,
      padding: EdgeInsets.fromLTRB(40, 20 + MediaQuery.of(context).padding.top, 40, 20),
      child: Column(
        children: [
          Divider(color: Theme.of(context).primaryColor,),
          SizedBox(
            width: double.infinity,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.start,
              children: [
                const SizedBox(height: 20,),
                _buildUserID(context),
                const SizedBox(height: 20,),
                _buildSolvedCount(),
                const SizedBox(height: 20,),
                _buildMaxStreak(),
                const SizedBox(height: 20,),
                _buildClass(),
                const SizedBox(height: 20,),
                _buildTierAndRating(),
                const SizedBox(height: 40,),
              ],
            ),
          ),
          Divider(color: Theme.of(context).primaryColor,),
          SizedBox(
            width: double.infinity,
            child: Column(
              mainAxisAlignment: MainAxisAlignment.start,
              crossAxisAlignment: CrossAxisAlignment.center,
              children: [
                const SizedBox(height: 20,),
                const Text("추천 문제"),
                const SizedBox(height: 20,),
                Text("문제 번호: ${viewModel.problem?.problemId ?? ""}"),
                const SizedBox(height: 20,),
                Text("문제 레벨: ${viewModel.problem?.level ?? ""}"),
                const SizedBox(height: 40,),
              ],
            ),
          ),
          const Spacer(),
          const Image(image: AssetImage('images/lmnop_logo_banner.png')),
        ],
      ),
    );
  }

  Widget _buildUserID(BuildContext context) {
    return RichText(
      text: TextSpan(
        children: <TextSpan>[
          TextSpan(text: "Hi, ", style: TextStyle(color: Theme.of(context).primaryColorDark)),
          TextSpan(text: viewModel.user?.handle ?? "", style: TextStyle(color: Theme.of(context).primaryColor)),
        ],
        style: const TextStyle(fontSize: 30, fontWeight: FontWeight.w700),
      ),
    );
  }

  Widget _buildSolvedCount() {
    return RichText(
      text: TextSpan(
        children: <TextSpan>[
          const TextSpan(text: "총 문제 해결: ", style: TextStyle(fontWeight: FontWeight.w300)),
          TextSpan(text: "${viewModel.user?.solvedCount ?? ""}", style: const TextStyle(fontWeight: FontWeight.w700)),
          const TextSpan(text: "문제", style: TextStyle(fontWeight: FontWeight.w300)),
        ],
        style: const TextStyle(fontSize: 20, color: Colors.black),
      ),
    );
  }

  Widget _buildMaxStreak() {
    return RichText(
      text: TextSpan(
        children: <TextSpan>[
          const TextSpan(text: "최대 연속 문제 해결: ", style: TextStyle(fontWeight: FontWeight.w300)),
          TextSpan(text: "${viewModel.user?.maxStreak ?? ""}", style: const TextStyle(fontWeight: FontWeight.w700)),
          const TextSpan(text: "일", style: TextStyle(fontWeight: FontWeight.w300)),
        ],
        style: const TextStyle(fontSize: 20, color: Colors.black),
      ),
    );
  }

  Widget _buildTierAndRating() {
    return RichText(
      text: TextSpan(
        children: <TextSpan>[
          TextSpan(text: "${tierToTierName[viewModel.user?.tier] ?? ""} ${viewModel.user?.rating ?? ""}",
              style: tierEnumToTextStyle[tierToTierEnum[viewModel.user?.tier]]),
        ],
        style: const TextStyle(fontSize: 20, fontWeight: FontWeight.w700,),
      ),
    );
  }

  Widget _buildClass() {
    return RichText(
      text: TextSpan(
        children: <TextSpan>[
          TextSpan(text: "${classToClassName[viewModel.user?.classNumber] ?? ""} ${classDecorationToClassDecorationName[viewModel.user?.classDecoration] ?? ""}",
            style: classEnumToTextStyle[classToClassEnum[viewModel.user?.classNumber]]),
        ],
        style: const TextStyle(fontSize: 20, fontWeight: FontWeight.w700,),
      ),
    );
  }
}
