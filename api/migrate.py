from tortoise import Tortoise
from settings import settings
from auth.auth import get_password_hash
from models import Site, User
from utils.error import ignore_async_errors


@ignore_async_errors
async def init_data():
    # 初始化数据
    await Tortoise.init(
        config=settings.DATABASE_CONFIG
    )
    # 创建表
    await Tortoise.generate_schemas()
    # 如果没有数据，创建一个超级管理员
    if not await User.all().exists():
        print('Creating super user')
        await User.create_one(
            {"username": "admin", "password": get_password_hash("123zh0.0"), "nickname": "超级管理员", "status": 1})
    # 如果没有数据，创建一个默认站点
    if not await Site.all().exists():
        await Site.create_one({
            "title": 'zh导航',
            "icon": 'ion:logo-edge',
            "desc": 'zh导航',
            "keywords": 'zh导航',
            "color": '#104A84',
            "copyright": '',
            "footer": 'Copyright zh导航. All Rights Reserved.'
        })
    await Tortoise.close_connections()
