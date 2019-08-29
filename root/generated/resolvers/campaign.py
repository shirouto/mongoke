
from tartiflette import Resolver
from .support import strip_nones, connection_resolver, zip_pluck, select_keys
from operator import setitem
from funcy import select_keys

@Resolver('Query.campaign')
async def resolve_query_campaign(parent, args, ctx, info):
    where = strip_nones(args.get('where', {}))
    headers = ctx['request']['headers']
    jwt_payload = ctx['req'].jwt_payload # TODO i need to decode jwt_payload and set it in req in a middleware
    fields = []

    collection = ctx['db']['campaigns']
    x = collection.find_one(where)

    if ('messages' in x):
        x['_typename'] = 'MessageCampaign'
    
    elif ('posts' in x):
        x['_typename'] = 'PostCampaign'
    
    if fields:
        x = select_keys(lambda k: k in fields, x)
    return x