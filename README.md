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

### 이진탐색트리(Binary Search Tree)
이진 탐색 트리는 빠른 검색을 위해 많이 사용된다.  
문자열 탐색시에도 문자열을 숫자로 만들어 이진트리에 넣고 빠르게 찾을 수 있다.  
왼쪽 자식노드는 부모노드보다 작고, 오른쪽 자식노드는 부모노드보다 크다.  
- 최소신장트리  
  - 효과적인 트리가 되기 위해선 트리의 높이가 최소가 되어야 한다.
- AVL트리
  - AVL트리는 스스로 균형을 잡는 이진 탐색 트리다. 왼쪽과 오른쪽의 서브트리의 높이 차이를 1이하로 유지한다.


### 힙
- 최대힙 : 부모노드가 자식노드보다 항상 크거나 같은 트리
- 최소힙 : 부모노드가 자식노드보다 항상 작거나 같은 트리
최대값 또는 최소값을 찾기에 좋다.  
힙으로 우선순위 큐를 만들 수 있다.  
힙을 이용해 정렬하는 방법도 있다.(힙 정렬)
속도는 Push와 Pop의 경우 logN, 힙정렬의 경우 NlonN  
힙 구현을 위해 배열을 많이 사용한다.  
- N번째 Left : 2N + 1
- N번째 Right: 2N + 2
- N번째 노드의 부모: (N-1)/2

### Map(Dictionary)
key - value 쌍을 가지는 구조  
- Sorted Map(Ordered Map)  
  - 맵의 key를 이진 트리로 관리할 수 있다
  - 속도는 logN
- Hash Map
  - key를 해시함수에 넣어 반한되는 인덱스로 value배열에서 key에 해당하는 값을 찾을 수 있다
  - 속도는 상수 시간  
  - 장점 : 조회, 추가, 삭제연산이 빠르다. (O(1))
  - 단점 : 정렬이 어렵다. ==> sorted map을 사용하면 되지만 성능은 O(logN)으로 느리다.

#### Hash?
  - 출력값의 범위가 있다.
  - 같은 입력이면 같은 출력이 나온다.
  - 다른 입력이면 보통의 경우 다른 출력이 나온다. (항상 만족하지는 않음)
  - 보통 Modular를 많이 쓴다. Mod N의 경우, 어떤 수를 넣어도 0~N-1까지의 정수가 나온다.
  ```
  Mod 연산은 One way function이다.
  x Mod N = M 일때,
  x는 무한대로 존재한다.
  따라서 해시의 결과물인 M값을 알아도 x가 무한히 존재하르모 x를 유추할 수가 없다.

  => 공개키 암호화, checksum, 블록체인 등에서 많이 쓰인다.
  ```

#### 롤링 해시
```
i번째 해시값 = (i-1번째 해시 X 상수 + i번째 문자) % 상수(소수)
0번째 해시값 = 문자 % 상수(소수)
문자열 길이만큼 반복해서 해시 계산을 했을 때, 최종적으로 나오는 값이 롤링 해시의 값이다.
```

### Thread
CPU가 수행해야할 작업 목록 - 쓰레드  
CPU는 쓰레드를 바꿔가며 작업 수행(멀티 태스킹)
  - 쓰레드 교체를 컨텍스트 스위칭이라 한다.
  - 컨텍스트 스위칭이 빈번하게 일어나면 성능이 저하된다.
OS는 CPU가 어떤 쓰레드를 가지고 작업할 지 총관리  
Go는 OS가 관리하는 쓰레드를 포장해서 Go Thread를 만듦.
  - OS 쓰레드를 최소한으로 사용하고 이를 포장한 후 잘게 쪼개서 사용
    - NM쓰레드: 하나의 OS 쓰레드에 여러개의 Go Thread가 들어갈 수 있다.
==> 컨텍스트 스위칭을 최소화하기 위해 Go는 CPU개수만큼만 OS Thread를 가져와 이를 잘게 쪼개서 사용하는 것이다.  
자바나 C#은 쓰레드풀을 이용



