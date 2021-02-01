# Go 학습(Tucker 의 프로그래밍 강좌)
https://youtu.be/SH32PgYGYRY

### 노트

### 23 -  Garbage Collector
- C언어는 메모리를 스택메모리와 힙메모리 구조로 나눈다.
- 변수를 선언하면 스택메모리에 쌓인다.
- 프로그래머가 힙 메모리의 용량을 할당 받을 수 있다.(malloc() 함수 사용)
```
p = malloc(100); // p는 할당받은 메모리의 주소를 가진다.

free(p); // C언어에서는 할당받은 메모리가 필요없어지면 지워줘야 한다.

// 버그 또는 실수로 메모리를 지우지 않은 경우 메모리에 쓰레기가 쌓인다. => memory leak => 메모리 공간이 부족해져 프로그램이 죽을 수 있다.

if (...) {
  int a; // int a는 스택메모리에 쌓인다. if문을 벗어나면 자동으로 스택메모리에서 지운다.
  ...
}
```

Garbage Collector는 메모리에 쌓인 쓰레기를 자동으로 치워준다.
1. 쓸모 없어지면 쓰레기다.
```
func add() {
  var a int
  a = 3
} // 함수를 빠져나가면 a가 쓸모 없어지므로 쓰레기가 된다.
GC가 쓰레기가 된 a를 치워준다.
```
#### Reference count
: 변수를 참조하고 있는 횟수를 기록한다. 0이 되면 쓰레기가 된다.
```
func add() *int {
  var a int // a의 ref count = 1
  a = 3
  var p *int // p의 ref count = 1
  p = &a // a의 ref count = 2

  return p
}

v:= add() // v의 ref count = 1, a의 ref count = 1
// p의 ref count = 0. p는 쓰레기가 됨.
```
C언어에서는 위의 예제의 경우 add()함수를 벗어나면 스택메모리에서 a와 p가 사라지므로 v는 빈 메모리를 참조하게 된다.(댕글링)
=> access violation으로 프로그램이 강제 종료된다.

#### 외딴섬
```
a -> b -> c-> a
// 이 경우 a와 b와 c는 모두 레퍼런스 카운트가 1이지만 다른곳에서 쓰이질 않으므로 GC가 치워준다.
```

### 24 - Slice
동적 배열을 어떻게 만들까?
- 새 크기의 배열을 만들고 기존 배열의 값들을 복사해온 뒤, 기존 배열을 삭제한다.
- 새 크기의 배열을 만들 때 필요한 길이보다 여유분을 확보해서 만든다.(2배씩 늘림)
- 동적 배열은 고정길이 배열을 포인트한다.
Slice도 하나의 struct로 보면 된다.
- Slice의 property?
  - pointer = 시작주소
  - len =  엘리먼트의 개수
  - cap = 배열이 저장할 수 있는 최대 엘리먼트 개수

#### 기본자료구조
배열, 링크드 리스트
이 두가지 기본 구조로 스택, 큐, 트리, 그래프, 힙, 맵 등을 만들 수 있다.

#### slice vs linked list
```
slice는 삽입 삭제 시 시간복잡도 O(n)
double linked list는 O(1)

하지만 조회의 경우 slice는 O(1)
linked list는 O(n) : Random Access

slice는 연속된 메모리 구조이므로 캐시에 올리기 좋음.(속도 향상)
반면 linked list는 메모리상에 흩어져 있으므로 캐시 미스가 자주 발생할 수 있음.
```

### DFS
깊이 우선 탐색은 최단거리를 구하는데 쓰이는 다익스트라 알고리즘에서도 사용된다.  
DFS는 재귀호출 또는 스택으로 구현할 수 있다.

### BFS
BFS는 큐를 이용해 구현한다.

### 이진탐색트리
이진 탐색 트리는 빠른 검색을 위해 많이 사용된다.  
문자열 탐색시에도 문자열을 숫자로 만들어 이진트리에 넣고 빠르게 찾을 수 있다.
