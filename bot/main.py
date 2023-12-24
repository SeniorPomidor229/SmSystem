import telebot
import requests
import pandas as pd

TOKEN = '6520254439:AAFhzmgkQiQBIn6TOIxAXG_nc6tAJEUWu1U'
bot = telebot.TeleBot(TOKEN)

@bot.message_handler(commands=['start'])
def start_command(message):
    bot.reply_to(message, "Привет! Чтобы получить данные и экспортировать их в файл XLSX, используйте команду /export")

@bot.message_handler(commands=['export'])
def export_data(message):
    try:
        base_url = "http://app:8000/api/sert/"
        date_text = message.text.split()[1] if len(message.text.split()) > 1 else ""

        if date_text:
            api_url = base_url + date_text
        else:
            api_url = base_url

        response = requests.get(api_url)

        if response.status_code == 200:
            data = response.json()

            if "detail" in data:
                results = data["detail"]

                df = pd.DataFrame(results)

                file_name = "exported_data.xlsx"
                df.to_excel(file_name, index=False)

                with open(file_name, 'rb') as file:
                    bot.send_document(message.chat.id, file, caption="Ваши данные в файле XLSX")
        elif response.status_code == 404:
            bot.reply_to(message, "Данных для указанной даты нет")
        else:
            bot.reply_to(message, "Произошла ошибка при получении данных")

    except Exception as e:
        print("Произошла ошибка:", e)
        bot.reply_to(message, "Произошла ошибка при экспорте данных")

bot.polling()
