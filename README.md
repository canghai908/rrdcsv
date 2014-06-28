#读取rrd文件的内容，并导出到csv文件#
最近研究rrd文件结构，想用工具将rrd文件内的数据导出到一个csv文件，以便统计和以后插件。最近看了点golang的基础，于是就边学边查，写了个小程序，希望对大家有用。
##工具要求##
* rrdtool 1.4版本以上，1.3以下版本不能用（rrdtool -v，可查看rrdtool版本）
* linux系统，win版本和freebsd版本目前还没有编译出来
##功能概括##
* 读取指定rrd文件最后288条数据到csv文件（也就是一天，5分钟一次）
* 导出文件名为执行命令时间向前86400秒（24小时）
* 导出文件内容为时间，流出，流入，单位Mbps
##使用方法##
 <pre> wget http://dl.cactifans.org/tools/rrd_csv.x32.tar.gz
tar zxvf rrd_csv.x32.tar.gz
chmod +x rrd_csv/fetch<code>
2.移动并编辑conf.json文件
移动conf.json到/etc目录下
<pre>cp rrd_csv/conf.json /etcconf.json<code>
文件内容
2014-04-23_165103
deviename为导出csv文件前缀
dbfile为需要导出的rrd文件名以及路径
3. 执行导出
./fetch
4.查看导出结果。
执行后没有报错，可在当前目录下查看，已生成以时间命名的cvs文件，形如

2014-04-23_181351

定时导出一天数据

1.下载导出工具到任意目录(我以opt目录为例）
cd /opt
wget http://dl.cactifans.org/tools/rrd_csv.x32.tar.gz
tar zxvf rrd_csv.x32.tar.gz
cp rrd_csv/conf.json /etc
chmod +x rrd_csv/fetch
2.添加计划任务
echo "6 12 * * * root /opt/rrd_csv/fetch">> /etc/crontab
注：建议每天凌晨12点6分导出，这样就可以导出前一天的完整数据
本人第一个golang程序，问题多多，以后会不断完善，希望大家多多指教啊。