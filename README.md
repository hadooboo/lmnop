<img height="160" src="https://user-images.githubusercontent.com/32262002/187039878-515910d9-8696-42a7-9634-9b1b71f7a52a.png">

## Demo

[Android APK](https://github.com/hadooboo/lmnop/files/9438328/app-release.apk.zip)

## Project Summary

[백준 온라인 저지](https://www.acmicpc.net/) 사이트에서 다음으로 어떤 문제를 풀지 자동으로 골라주기 위해 만들기 시작한 앱

귀찮은 로그인 과정 없이 BOJ 사이트의 닉네임 입력만으로 문제 선정

문제를 선정하는 기준 ->

> 내가 푼 모든 문제를 레벨 순서로 내림차순 정렬한 후, 상위 50번 째 문제의 레벨을 적정 레벨로 정의

> 적정 레벨에 속한 문제 중 내가 아직 풀지 않았으면서 가장 많은 사람들이 푼 문제를 선정

안드로이드 및 iOS 환경에서 모두 사용 가능!

## Used Stacks

### Backend

Go(Gin, Resty, Zap)

### Frontend

Flutter

### Deployment

Heroku

## 프로젝트 회고

<details><summary>외부 API 사용하여 매시업</summary><p>

전형적인 Frontend Client + Backend Server + Database 구조가 아닌 Backend Server가 외부 API의 클라이언트가 되는 구조로 만들었다. 내가 푼 문제에 대한 정보는 결국 BOJ 서비스에서 관리하고 있을 것이고 그 데이터를 가져다쓰는 입장이기 때문이다. BOJ 서비스에서 공식적으로 API를 제공하지는 않기 때문에 [solved.ac](https://solved.ac/) 사이트의 API를 이용하였다. 내가 데이터를 직접 관리하지 않아도 되는 것은 큰 장점이었다. 그러나 API의 주소나 응답 형식이 바뀔 경우 내 서비스가 작동하지 않게 될 수 있기 때문에 지속적인 모니터링은 필요할 듯 하다.

<p></details>

<details><summary>완벽하게 DDD 철학에 맞춘 설계</summary><p>

go로 만든 백엔드 서버의 코드를 살펴보면 [clean-architecture](https://github.com/wikibook/clean-architecture) 의 예시를 정확하게 따르고 있다. 이 레포지토리는 <만들면서 배우는 클린 아키텍처, 2021, 톰 홈버그>의 한국어판 예시이다. 이 프로젝트를 시작하게 된 이유가 위 책을 읽고 DDD 방법론에 맞춰 서버를 만들고 싶었던 것도 있다. DDD의 장점을 알 수 있었던 점이 책에서는 영속성 관리를 위해 데이터베이스를 사용하지만 이 프로젝트에서는 외부 API를 사용하는데, 이런 차이점에 구애받지 않고 out adapter를 구현할 수 있었던 것이다. 앞으로도 책의 내용을 내 것으로 만들기 위해 여러 사이드 프로젝트를 진행해봐야겠다.

추가로, application layer에 있는 [problem.go](https://github.com/hadooboo/lmnop/blob/main/server/application/problem.go) 파일에 문제 선정 로직이 작성되어 있는데, 이 역시 <Clean Code, 2009, 로버트 C. 마틴> 책의 내용을 보고 영감을 받아 만든 부분이다. 로직을 위한 전용 클래스를 만들고 내부 변수에 필요한 변수들을 모두 정의해 둔 다음 메소드에는 파라미터 전달을 최소한으로 하고 최대한 자연어로 읽을 수 있게 구현하였다.

<p></details>

<details><summary>철저하려고 노력한 테스트 코드</summary><p>

프로젝트의 규모가 크진 않았지만 나름의 문제 선정 로직이 들어가기 때문에 테스트 코드를 철저히 만들어보자고 마음을 먹었다. [service_test.go](https://github.com/hadooboo/lmnop/blob/main/server/application/service_test.go) 파일을 보면 여러 개의 테스트 케이스를 정의해서 확인하려고 한 흔적을 볼 수 있다. 외부 API 응답에 대한 mocking을 만들어 application layer에서 로직에 빈틈이 없는지 확인하였다. 다시 보니 다음에는 테스트 단위를 더 작게 분리하고 케이스 각각에 대한 설명도 추가하면 좋을 것 같다.

<p></details>

<details><summary>Flutter를 사용하여 앱 만들기</summary><p>

데스크탑 환경에서는 solved.ac 사이트를 쓰는 것만으로도 충분하고, 언제 어디서든 다음 목표 문제를 보기 위해 클라이언트를 모바일 앱으로 만들기로 정했다. 특히 개인적으로 갤럭시 스마트폰에 아이패드를 사용하고 있어서 회사에서도 이용해 본 flutter를 이용해 멀티 플랫폼을 타겟팅하였다. 분명히 두 번 할 작업을 한 번의 개발만으로 끝낼 수 있어서 잘 선택했다고 생각한다. 특히 이번 프로젝트처럼 간단한 수준의 앱은 성능이나 기능 면에서 native SDK와 비교했을 때 문제가 없어 보인다.

<p></details>

<details><summary>Heroku를 이용한 API 서버 배포</summary><p>

실제로 앱을 플레이스토어나 앱스토어에 출시하지는 않았기 때문에 API 서버를 비용을 내고 실행하고 있는 것은 아깝다는 생각이 들었다. Heroku는 한 달에 최대 1000시간의 무료 사용 시간을 제공하고, 서버를 사용하지 않고 있다고 판단하면 알아서 잠자기 모드로 들어가기 때문에 내 상황에서는 충분히 사용 가능하다. go module을 포함하고 있는 레포지토리를 git을 통해 푸시하는 것만으로도 배포가 자동적으로 이뤄져 편리하였다.

<p></details>

<details><summary>서버리스로의 확장?</summary><p>

현재 프로그램의 구조를 보면 모든 곳이 stateless하게 작동할 수 있다. 이런 상황에서 굳이 서버를 계속해서 돌리고 있을 것이 아니라 메소드를 필요할 때만 호출하는 서버리스로 만든다면 비용도 절약하고 컴퓨팅 파워도 아낄 수 있을 것이라고 생각한다. Heroku로 서버를 실행해 두었기 때문에 서버리스의 단점이라고 할 수 있는 cold start가 어차피 지금도 발생하고 있다. 다음에 기회가 된다면 서버리스로의 확장도 해 보고 싶다.

<p></details>
