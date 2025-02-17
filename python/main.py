import structlog

logger = structlog.get_logger()

def main():
    logger.info("Hello World from Python!")


if __name__ == "__main__":
    main()
