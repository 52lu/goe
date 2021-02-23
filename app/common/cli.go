/**
 * @Author Mr.LiuQH
 * @Description 打印到控制台
 * @Date 2021/2/23 11:42 上午
 **/
package common

import (
	"github.com/olekukonko/tablewriter"
	"os"
)

type CliInfo struct {
	ConfigFile string
	Host       string
	Port       string
}

/**
 * @description: 打印启动信息
 * @user: Mr.LiuQH
 * @receiver cli CliInfo
 * @date 2021-02-23 11:58:02
 */
func (cli CliInfo) PrintRunMsg() {
	data := [][]string{
		{"Config", cli.ConfigFile},
		{"Host", cli.Host},
		{"Port", cli.Port},
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"key","Value"})
	table.SetColWidth(800)
	// 设置颜色
	table.SetColumnColor(
		tablewriter.Colors{tablewriter.FgHiGreenColor},
		tablewriter.Colors{tablewriter.FgHiGreenColor},
	)
	// 设置居中显示
	table.SetAlignment(tablewriter.ALIGN_LEFT)
	// 填充数据
	table.AppendBulk(data)
	table.Render()
}
