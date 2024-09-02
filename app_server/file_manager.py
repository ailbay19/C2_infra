import secrets
from queue import Queue

def random_string(length = 16):
    return secrets.token_hex(length)

class FileManager:
    
    def __init__(self):
        self.queue = Queue()
    
    def get(self):
        if self.queue.empty():
            return None
        res = self.queue.get()
        
        # TODO: FOR DEVELOPMENT
        self.queue.put(res)
        
        return res
    
    def put(self, file):
        res = {}
        
        res['filename'] = random_string()
        res['file'] = file
        
        self.queue.put(res)
        
        return res