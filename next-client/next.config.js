require('dotenv').config();

/** @type {import('next').NextConfig} */
const nextConfig = {
  reactStrictMode: true,
  async rewrites() {
		return [
			{
				source: '/api/:path*',
				destination: `${process.env.API_BASE_URL}/:path*`,
			},
		]
	},
  env: {
    GOOGLE_CLIENT_ID: process.env.GOOGLE_CLIENT_ID
  },
};

module.exports = nextConfig;
