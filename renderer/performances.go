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

type Performances struct {
	ProgramCPU      string
	ProgramRAM      string
	ProgramGCCycles string
	SystemCPU       string
	SystemTotalRAM  string
	SystemUsedRAM   string
	TotalCores      int
}

const OFFSET float32 = 10
const WIDTH float32 = 400
const HEIGHT float32 = 200
const FONTSIZE float32 = 24
const PADDING float32 = 10

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

func (p *Performances) RenderPerformanceDisplay(style st.WindowStyle, state t.ProgramState) {
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
	performanceText := fmt.Sprintf("Program:\nCPU: %s%%\nRAM: %sMiB\n\nSystem:\nCPU: %s%%\nRAM: %sMiB / %sMiB",
		p.ProgramCPU, p.ProgramRAM, p.SystemCPU, p.SystemUsedRAM, p.SystemTotalRAM)
	rl.DrawTextEx(style.Font, performanceText, textPos, FONTSIZE, style.FontSpacing, TEXT_COLOR)
}

func (p *Performances) CalculatePerformanceDisplay(ctx context.Context) {
	var m runtime.MemStats
	pid := fmt.Sprintf("%d", os.Getpid())

	for {
		select {
		case <-ctx.Done():
			return
		default:
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
			time.Sleep(time.Second)
		}
	}
}

func getProgramCPUUsage(pid string) string {
	out, err := exec.Command("ps", "-p", pid, "-o", "%cpu=").Output()
	if err != nil {
		return "Error"
	}
	return strings.TrimSpace(string(out))
}
