import asyncio
import aiogram
import requests
import pandas as pd

bot = aiogram.Bot(token='6520254439:AAFhzmgkQiQBIn6TOIxAXG_nc6tAJEUWu1U')
dp = aiogram.Dispatcher(bot)

@dp.message_handler(commands=['start'])
async def start_command(message: aiogram.types.Message):
    await message.reply("Привет! Чтобы получить данные и экспортировать их в файл XLSX, используйте команду /export")

@dp.message_handler(commands=['export'])
async def export_data(message: aiogram.types.Message):
    try:
        base_url = "http://127.0.0.1:8000/api/sert/"
        date_text = message.get_args()

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
                    await bot.send_document(message.chat.id, file, caption="Ваши данные в файле XLSX")
        elif response.status_code == 404:
            await message.reply("Данных для указанной даты нет")
        else:
            await message.reply("Произошла ошибка при получении данных")

    except Exception as e:
        print("Произошла ошибка:", e)
        await message.reply("Произошла ошибка при экспорте данных")


if __name__ == '__main__':
    aiogram.executor.start_polling(dp, skip_updates=True)
