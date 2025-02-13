import structlog
from rich.console import Console

logger = structlog.get_logger()
console = Console()


def add_numbers(a: int, b: int) -> int:
    return a + b


def main():
    # Log a hello world message
    logger.info("Hello, World from Python!")
    logger.warning("This is a warning.")
    logger.error("This is an error.")

    try:
        # This will cause a mypy error because we're passing a string instead of an int
        result = add_numbers(5, "10")
        logger.info(f"The result is {result}")

        # Use rich to print a message to the console
    except Exception:
        console.print_exception(show_locals=True)


if __name__ == "__main__":
    main()
