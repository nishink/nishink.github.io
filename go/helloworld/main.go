// Description: Ebitenを使ってHello, World!を表示する
// https://ebitengine.org/ja/
package main

// パッケージのインポート
// log: ログ出力
// ebiten: ゲームエンジン
// ebitenutil: デバッグ用の関数
import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

// Game構造体
// ゲームの状態を保持する
type Game struct{}

// Update関数
// ゲームの状態を更新する
func (g *Game) Update() error {
	return nil
}

// Draw関数
// ゲーム画面を描画する
func (g *Game) Draw(screen *ebiten.Image) {
	// デバッグ用の関数を使ってHello, World!を表示する
	ebitenutil.DebugPrint(screen, "Hello, World!")
}

// Layout関数
// ウィンドウのサイズを設定する
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

// main関数
// ウィンドウのサイズを設定し、ゲームを実行する
func main() {
	// ウィンドウのサイズを設定
	ebiten.SetWindowSize(640, 480)
	// ウィンドウのタイトルを設定
	ebiten.SetWindowTitle("Hello, World!")
	// ゲームを実行
	if err := ebiten.RunGame(&Game{}); err != nil {
		// エラーが発生した場合はログに出力
		log.Fatal(err)
	}
}