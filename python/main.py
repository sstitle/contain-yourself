import coloredlogs, logging

def main():
    # Create a logger object
    logger = logging.getLogger(__name__)

    # Set up coloredlogs
    coloredlogs.install(level='DEBUG', logger=logger)

    # Log a hello world message
    logger.info("Hello, World from Python!")

if __name__ == "__main__":
    main()

