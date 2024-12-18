package renderer

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"time"

	st "github.com/basileb/kenzan/settings"
	t "github.com/basileb/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/mem"
)

type Performance struct {
	ProgramCPU      string
	ProgramRAM      string
	ProgramGCCycles string
	SystemCPU       string
	SystemTotalRAM  string
	SystemUsedRAM   string
	TotalCores      int
}

type Performances = []Performance

const OFFSET float32 = 10
const WIDTH float32 = 400

const HEIGHT float32 = 200
const FONTSIZE float32 = 24
const PADDING float32 = 10

const MEASURE_NB int32 = 30
const GRAPH_RATIO float32 = 1

var (
	darkGray      = rl.Color{30, 30, 30, 216} // #1e1e1e - Modern dark gray
	nearBlack     = rl.Color{18, 18, 18, 216} // #121212 - Spotify-like black
	pureBlack     = rl.Color{0, 0, 0, 216}    // #000000 - Pure black
	mutedBlue     = rl.Color{44, 62, 80, 216} // #2c3e50 - Soft muted blue
	terminalGreen = rl.Color{0, 59, 0, 216}   // #003b00 - Retro terminal green
	charcoalGray  = rl.Color{45, 45, 45, 216} // #2d2d2d - High contrast modern gray
)

var (
	neonGreen    = rl.Color{57, 255, 20, 255}   // #39ff14 - Neon green
	cyberCyan    = rl.Color{0, 255, 255, 255}   // #00ffff - Cyber cyan
	amberCRT     = rl.Color{255, 191, 0, 255}   // #ffbf00 - Vintage amber CRT
	pastelYellow = rl.Color{245, 222, 179, 255} // #f5deb3 - Soft pastel yellow
	brightRed    = rl.Color{231, 76, 60, 255}   // #e74c3c - High contrast red
	brightBlue   = rl.Color{78, 154, 241, 255}  // #4e9af1 - High contrast blue
	offWhite     = rl.Color{236, 239, 241, 255} // #ECEFF1 - Clean off-white
)

var TEXT_COLOR rl.Color = cyberCyan
var BG_COLOR rl.Color = mutedBlue

func RenderPerformanceDisplay(perfs Performances, style st.WindowStyle, state t.ProgramState) {
	rec := rl.Rectangle{
		X:      state.ViewPortSize.X - WIDTH - OFFSET,
		Y:      OFFSET,
		Width:  WIDTH,
		Height: HEIGHT,
	}
	rl.DrawRectangleRounded(rec, 0.08, 4, BG_COLOR)

	textPos := rl.Vector2{
		X: rec.X + PADDING,
		Y: rec.Y + PADDING,
	}
	if len(perfs) > 0 {
		i := len(perfs) - 1
		performanceText := fmt.Sprintf("Program:\nCPU: %s%%\nRAM: %sMiB\n\nSystem:\nCPU: %s%%\nRAM: %sMiB / %sMiB",
			perfs[i].ProgramCPU, perfs[i].ProgramRAM, perfs[i].SystemCPU, perfs[i].SystemUsedRAM, perfs[i].SystemTotalRAM)
		rl.DrawTextEx(style.Font, performanceText, textPos, FONTSIZE, style.FontSpacing, TEXT_COLOR)
		drawGraph(perfs, style, state)
	}
}

func drawGraph(perfs Performances, style st.WindowStyle, state t.ProgramState) {
	rec := rl.Rectangle{
		X:      state.ViewPortSize.X - WIDTH - OFFSET,
		Y:      HEIGHT + 2*OFFSET,
		Width:  WIDTH,
		Height: GRAPH_RATIO * HEIGHT,
	}
	rl.DrawRectangleRounded(rec, 0.08, 4, BG_COLOR)

	// Title
	title := "Program CPU usage"
	const TITLE_FONTSIZE float32 = 18
	titleSize := rl.MeasureTextEx(style.Font, title, TITLE_FONTSIZE, style.FontSpacing)
	titlePos := rl.Vector2{
		X: rec.X + WIDTH/2 - titleSize.X/2,
		Y: rec.Y + PADDING,
	}
	rl.DrawTextEx(style.Font, title, titlePos, TITLE_FONTSIZE, style.FontSpacing, offWhite)

	graphHeight := int32(rec.Height - 2*PADDING)
	graphWidth := int32(rec.Width - 2*PADDING)

	xIntervals := graphWidth / MEASURE_NB
	prevV := float64(0)
	if len(perfs)-int(MEASURE_NB) > 0 {
		prevV, _ = strconv.ParseFloat(perfs[int(len(perfs)-int(MEASURE_NB))].ProgramCPU, 32)
	}
	for i := range MEASURE_NB {
		if len(perfs)-int((1+i)) <= 0 {
			break
		}

		currentV, _ := strconv.ParseFloat(perfs[len(perfs)-int((1+i))].ProgramCPU, 32)
		currentYPos := (currentV / 100) * float64(graphHeight)

		prevYPos := (prevV / 100) * float64(graphHeight)
		rl.DrawLine(int32(rec.X+2*PADDING)+xIntervals*(i), int32(rec.Y+GRAPH_RATIO*HEIGHT-2*PADDING)-int32(prevYPos),
			int32(rec.X+2*PADDING)+xIntervals*(i+1), int32(rec.Y+GRAPH_RATIO*HEIGHT-2*PADDING)-int32(currentYPos), TEXT_COLOR)

		if len(perfs)-int(2+i) > 0 {
			prevV = currentV
		}
	}

	// X axis
	rl.DrawRectangle(int32(rec.X+2*PADDING), int32(rec.Y+GRAPH_RATIO*HEIGHT-2*PADDING), int32(WIDTH-4*PADDING), 1, offWhite)

	// Y axis
	rl.DrawRectangle(int32(rec.X+2*PADDING), int32(rec.Y+2*PADDING), 1, int32(GRAPH_RATIO*HEIGHT-4*PADDING), offWhite)

	// 100% text
	const TEXT_FONTSIZE float32 = 12
	textPos := rl.Vector2{
		X: rec.X + 2,
		Y: rec.Y + PADDING - 4,
	}
	rl.DrawTextEx(style.Font, "100%", textPos, TEXT_FONTSIZE, style.FontSpacing, offWhite)

	// 0% text
	textPos = rl.Vector2{
		X: rec.X + 2,
		Y: rec.Y + GRAPH_RATIO*HEIGHT - 4*PADDING + 4,
	}
	rl.DrawTextEx(style.Font, "0%", textPos, TEXT_FONTSIZE, style.FontSpacing, offWhite)
}

func GetPerformances(ctx context.Context, perf *Performances) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			var p Performance
			p.CalculatePerformanceDisplay()
			*perf = append(*perf, p)
			time.Sleep(time.Millisecond * 10)
		}
	}

}

func (p *Performance) CalculatePerformanceDisplay() {
	var m runtime.MemStats
	pid := fmt.Sprintf("%d", os.Getpid())

	p.TotalCores = runtime.NumCPU()

	runtime.ReadMemStats(&m)
	programCPU := getProgramCPUUsage(pid)
	programCPUNum, _ := strconv.ParseFloat(programCPU, 32)
	programCPU = fmt.Sprintf("%.2f", programCPUNum/float64(p.TotalCores))
	vmStat, _ := mem.VirtualMemory()
	cpuPercent, _ := cpu.Percent(time.Second, false)

	p.ProgramCPU = programCPU
	p.ProgramRAM = fmt.Sprintf("%v", m.Alloc/1024/1024)
	p.ProgramGCCycles = fmt.Sprintf("%v", m.NumGC)
	p.SystemCPU = fmt.Sprintf("%.2f%%", cpuPercent[0])
	p.SystemUsedRAM = fmt.Sprintf("%v", vmStat.Used/1024/1024)
	p.SystemTotalRAM = fmt.Sprintf("%v", vmStat.Total/1024/1024)
}

func getProgramCPUUsage(pid string) string {
	out, err := exec.Command("ps", "-p", pid, "-o", "%cpu=").Output()
	if err != nil {
		return "Error"
	}
	return strings.TrimSpace(string(out))
}

func findHighIndexes(performanceSlice Performances) int {
	if len(performanceSlice) == 0 {
		return -1
	}

	highestIndex := 0

	for i := 1; i < len(performanceSlice); i++ {
		// Convert ProgramCPU strings to float64 for comparison
		currentCPU, _ := strconv.ParseFloat(performanceSlice[i].ProgramCPU, 64)
		highestCPU, _ := strconv.ParseFloat(performanceSlice[highestIndex].ProgramCPU, 64)

		if currentCPU > highestCPU {
			highestIndex = i
		}
	}

	return highestIndex
}
