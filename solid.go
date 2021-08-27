/*  SOLID 원칙 기반 Beyond OOP 구현 해커톤
    부산 센탑 goBlock 해커톤 특별 세트 메뉴 제작
	1원칙 : 단일 책임 원칙
	2원칙 : 개방-폐쇄 원칙
	3원칙 : 리스코프 치환 원칙
	4원칙 : 인터페이스 분리 원칙
	5원칙 : 의존관계 역전 원칙
*/

package main

import (
	"fmt"
)

// ------- 모둠회세트 메뉴 인터페이스 -------------//
// 인터페이스 1 - 부산 회
type Dish interface {
	String() string
}

// 인터페이스 2 - 회 뜨기
type Slice interface {
	getonedish() Dish
}

type Sashimi struct {
	val string
}

//  -------- Fillet 메소드 작성 ---------- //
func (s *Sashimi) Fillet(Slice Slice) {
	pickit := Slice.getonedish()
	s.val += pickit.String()
}

func (s *Sashimi) String() string {
	return "(( 멍게 " + s.val
}

// ----- 손질된 회를 접시에 올립니다 ------- //

// 광어를 회 떠요
type FlatfishSlice struct {
}

func (j *FlatfishSlice) getonedish() Dish {
	return &PieceofFlatfishSlice{}
}

type PieceofFlatfishSlice struct {
}

func (s *PieceofFlatfishSlice) String() string {
	return " + 광어 한 마리 "
}

// 우럭을 회 떠요
type RockfishSlice struct {
}

func (j *RockfishSlice) getonedish() Dish {
	return &PieceofRockfishSlice{}
}

type PieceofRockfishSlice struct {
}

func (s *PieceofRockfishSlice) String() string {
	return " + 우럭 한 마리 "
}

// 도미를 회 떠요
type BreamSlice struct {
}

func (j *BreamSlice) getonedish() Dish {
	return &PieceofBreamSlice{}
}

type PieceofBreamSlice struct {
}

func (s *PieceofBreamSlice) String() string {
	return " + 도미 한 마리"
}

// 우럭을 회 떠요
type BassSlice struct {
}

func (j *BassSlice) getonedish() Dish {
	return &PieceofBassSlice{}
}

type PieceofBassSlice struct {
}

func (s *PieceofBassSlice) String() string {
	return " + 농어 한 마리"
}

// ------- 서비스 메뉴 인터페이스 -------------//
// 인터페이스 3 - 서비스
type Service interface {
	String() string
}

// 인터페이스 4 - 사장님의 마음이 담긴 무료 서비스
type free interface {
	HeartofOwner() Service
}

type MarketService struct {
	val string
}

// ------ pickService 메소드 작성 ------- //
func (f *MarketService) pickService(Free free) {
	serve := Free.HeartofOwner()
	f.val += serve.String()
}

func (f *MarketService) String() string {
	return f.val
}

// 서비스 1번 대구탕 추가

type SpicyFishStew struct {
}

func (f *SpicyFishStew) HeartofOwner() Service {
	return &ServiceSpicyFishStew{}
}

type ServiceSpicyFishStew struct {
}

func (s *ServiceSpicyFishStew) String() string {
	return "생선머리 넣고 끓인 대구탕"
}

// 서비스 2번 콜라 추가

type Cola struct {
}

func (f *Cola) HeartofOwner() Service {
	return &MarketPepsiCola{}
}

type MarketPepsiCola struct {
}

func (s *MarketPepsiCola) String() string {
	return " + 업소용 콜라 1.5L"
}

func main() {
	// 회와 서비스 구조체를 불러옵니다.
	Sashimi := &Sashimi{}
	MarketService1 := &MarketService{}

	// 개별 횟감과 개별 서비스의 구조체를 불러옵니다.
	Slice := &BassSlice{}
	Slice1 := &FlatfishSlice{}
	Slice2 := &BreamSlice{}
	Slice3 := &RockfishSlice{}

	Service00 := &SpicyFishStew{}
	Service01 := &Cola{}

	// ----- 각 인터페이스에 해당하는 메소드를 불러옵니다 ------- //
	Sashimi.Fillet(Slice)
	Sashimi.Fillet(Slice1)
	Sashimi.Fillet(Slice2)
	Sashimi.Fillet(Slice3)

	MarketService1.pickService(Service00)
	MarketService1.pickService(Service01)

	// 완성!
	fmt.Println(" *:･｡,☆ﾟ’･:*:･*:･｡,☆ﾟ’･:*:･｡환    영｡･:*:･ﾟ’☆,｡･:**:･｡,☆ﾟ’･:*:･ ")
	fmt.Println("----------------부산 센탑 goBlock 해커톤 특별메뉴-----------------")
	fmt.Println("------------------- goBlock 부산 센텀시티점 ------------------")
	fmt.Println("")
	fmt.Println("(( ~~~~~~~~~~~~~~~~~~~~~~~~~ 천사채 ~~~~~~~~~~~~~~~~~~~~~~~~))")
	fmt.Println("(( ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~))")
	fmt.Println(Sashimi, "))")
	fmt.Println("(( ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~))")
	fmt.Println("(( ~~~~~~~~~~~~~~~~~~~~~~~~~ 천사채 ~~~~~~~~~~~~~~~~~~~~~~~~~))")
	fmt.Println("")
	fmt.Println("서비스 메뉴 : ", MarketService1, " 나왔습니다!")
	fmt.Println("")
	fmt.Println("-------------- 주문하신 부산 센탑 특별 모듬회 나왔습니다 ----------------")
	fmt.Println("----------------- 서비스 더 필요하시면 말씀해주세요 -------------------")
	fmt.Println("*:･｡,☆ﾟ’･:*:･*:･｡,☆ﾟ’･:*:･｡주문 감사합니다｡･:*:･ﾟ’☆,｡･:**:･｡,☆ﾟ’･:*:･")
}

/// 요구사항에 맞는 .. 답을 정해줬으니 어쨋든 답을 만들어라. (고객이 원한 요구사항...)
