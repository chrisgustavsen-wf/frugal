class LanguageBase(object):
    """Language update implementations must implement LanguageBase."""

    def update_frugal(self, version):
        """Update the frugal version."""
        raise Exception('update_frugal not implemented')
