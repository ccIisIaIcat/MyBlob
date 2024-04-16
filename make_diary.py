from bs4 import BeautifulSoup
import datetime
import pandas as pd
from bs4 import NavigableString

name = "diary_"+datetime.datetime.now().strftime("%Y-%m-%d")
path = "./static/article_html/"+name+".html"
diary_content = []

with open("./diary_demo.txt", "r", encoding="utf-8") as file:
    # 读取文件内容，并将每一行作为一个字符串放入列表中
    lines = file.readlines()
    # 输出列表中的每一行内容
    for line in lines:
        diary_content.append(line.strip())

# 读取HTML文件
with open("./static/article_html/diary_demo.html", "r", encoding="utf-8") as file:
    html_content = file.read()

# 使用Beautiful Soup解析HTML
soup = BeautifulSoup(html_content, "html.parser")

# 找到所有 class 为 "contt" 的 div 标签
contt_divs = soup.find_all("div", class_="contt")

# 遍历每个 div 标签，添加新的 <p> 标签
for div in contt_divs:
    for v in diary_content:
        new_paragraph1 = soup.new_tag("p")
        new_paragraph1.append(NavigableString("      "))
        new_paragraph1.append(NavigableString(v))
        # 将新的 <p> 标签插入到当前 div 标签中
        div.append(new_paragraph1)


# 修改路径
df = pd.read_csv("./data/Diary_index/index.csv",header=None)
df.loc[len(df)] = [name,name+".html"]
print(df)
df.to_csv("./data/Diary_index/index.csv",header=False, index=False)


# 保存修改后的HTML内容
with open("./static/article_html/"+name+".html", "w", encoding="utf-8") as file:
    file.write(str(soup))