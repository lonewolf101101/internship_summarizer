import { createServer } from 'https';
import { readFileSync } from 'fs';
import { resolve } from 'path';
import { loadNuxt, build } from 'nuxt';

const httpsOptions = {
  key: readFileSync(resolve(__dirname, 'localhost.key')),
  cert: readFileSync(resolve(__dirname, 'localhost.crt')),
};

async function start() {
  const nuxt = await loadNuxt('dev');
  
  const server = createServer(httpsOptions, nuxt.render);
  
  server.listen(3000, () => {
    console.log('HTTPS Server running at https://localhost:3000');
  });

  // Build in development mode
  if (process.env.NODE_ENV !== 'production') {
    await build(nuxt);
  }
}

start();
